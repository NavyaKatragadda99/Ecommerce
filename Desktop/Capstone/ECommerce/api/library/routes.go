package api

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(db *sql.DB) {

	http.HandleFunc("/add", AddHandler(db))
	http.HandleFunc("/update", UpdateHandler(db))
	http.HandleFunc("/delete", DeleteHandler(db))
	http.HandleFunc("/list", ListHandler(db))
	// http.HandleFunc("/list", listHandler(db))

	// http.HandleFunc("/register", registerHandler(db))
	// http.HandleFunc("/login", loginHandler(db))
}

// func listProducts(db *sql.DB) []Product {
// 	rows, err := db.Query("SELECT id, name, price FROM products")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	var products []Product
// 	for rows.Next() {
// 		var p Product
// 		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
// 			log.Fatal(err)
// 		}
// 		products = append(products, p)
// 	}
// 	return products
// }
