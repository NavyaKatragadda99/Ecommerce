package dataservice

import (
	"database/sql"
	"encoding/json"
	"example/Ecommerce/model"
	"fmt"
	"net/http"
)

// AuthenticatorData interface defines methods for user authentication and authorization
type AuthenticatorData interface {
	AddUser(authenticator model.Authenticator) error
	GetUser(username string) (model.Authenticator, error)
	GetUserRoles(userID int) ([]string, error)
}

// authenticatorData is the concrete implementation of AuthenticatorData
type authenticatorData struct {
	db *sql.DB
}

// NewAuthenticatorData creates a new instance of authenticatorData
func NewAuthenticatorData(db *sql.DB) *authenticatorData {
	return &authenticatorData{db: db}
}

func UserExists(db *sql.DB, username string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE username=?)"
	var exists bool
	err := db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking user existence: %v", err)
	}
	return exists, nil
}

func AddUserQuery(db *sql.DB, w http.ResponseWriter, authenticator model.Authenticator) error {
	query := `INSERT INTO users(username, password, email, role) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, authenticator.Username, authenticator.Password, authenticator.Email, authenticator.Role)
	if err != nil {
		return fmt.Errorf("error adding user: %v", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(authenticator)

	return nil
}

func AuthenticateUser(db *sql.DB, username, password string) (int, error) {
	var userID int
	query := "SELECT user_id FROM users WHERE username = ? AND password = ?"
	err := db.QueryRow(query, username, password).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("authentication failed: user not found or invalid credentials")
		}
		return 0, fmt.Errorf("error authenticating user: %v", err)
	}
	return userID, nil
}

func GetUserRoles(db *sql.DB, userID int) ([]string, error) {
	query := `SELECT role FROM users WHERE user_id = ?`
	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user roles: %v", err)
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var role string
		if err := rows.Scan(&role); err != nil {
			return nil, fmt.Errorf("error scanning user roles: %v", err)
		}
		roles = append(roles, role)
	}

	return roles, nil
}



