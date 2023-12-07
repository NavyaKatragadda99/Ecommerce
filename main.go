package main

import (
	"database/sql"
	"example/Ecommerce/api"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to the database
	db, err := sql.Open("mysql", "root:24021199@tcp(127.0.0.1:3306)/ecommerce_db?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Register API routes
	api.RegisterRoutes(db)

	// Start the HTTP server
	log.Println("Starting the server...")
	http.ListenAndServe(":8080", nil)
}
