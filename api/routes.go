package api

import (
	"database/sql"
	"net/http"
)

// RegisterRoutes registers API routes
func RegisterRoutes(db *sql.DB) {
	http.HandleFunc("/signup", AddUser(db))
	http.HandleFunc("/signin", LoginUser(db))
	http.HandleFunc("/profile", RoleUser(db))
}


