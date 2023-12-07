package api

import (
	"database/sql"
	"errors"
	"example/Ecommerce/dataservice"
	"example/Ecommerce/model"
	"net/http"
)

func SignUpLogic(db *sql.DB, w http.ResponseWriter, authenticator model.Authenticator) error {
	// Check if the user already exists
	if exists, err := dataservice.UserExists(db, authenticator.Username); err != nil {
		return err
	} else if exists {
		http.Error(w, "user already exists", http.StatusBadRequest)
		return errors.New("user already exists")
	}

	// User does not exist, proceed with signup
	return dataservice.AddUserQuery(db, w, authenticator)
}

func SignInLogic(db *sql.DB, username, password string) (string, error) {
	// Authenticate user and get user ID
	userID, err := dataservice.AuthenticateUser(db, username, password)
	if err != nil {
		return "", err
	}

	// TODO: Generate and return a token (or session) for the authenticated user
	token, err := generateToken(userID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Example token generation function (replace with your actual implementation)
func generateToken(userID int) (string, error) {
	// Implement your token generation logic here (e.g., JWT)
	// Return the generated token
	return "example_token", nil
}

