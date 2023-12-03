package main

import (
	api "ECommerce/api/library"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Connect to MySQL database
	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/sys?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MySQL!")

	api.RegisterRoutes(db)

	log.Println("server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
