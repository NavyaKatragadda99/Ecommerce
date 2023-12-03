package dataservice

import (
	"ECommerce/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func AddProduct(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return err
	}

	query := `INSERT INTO products (Id, Name, Description, Price) VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, product.ID, product.Name, product.Description, product.Price)
	if err != nil {
		return err
	}
	fmt.Println("Product added successfully!")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
	return nil
}

func ListProduct(db *sql.DB) ([]model.Product, error) {
	rows, err := db.Query("SELECT Id, Name,Description, Price FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price); err != nil {
			log.Fatal(err)
		}
		products = append(products, p)
	}
	fmt.Println("Product Listed successfully!")
	fmt.Println(products)
	return products, err
}

func UpdateProduct(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return err
	}

	query := `UPDATE products SET name = ?, Description= ? WHERE Id = ?`
	_, err := db.Exec(query, product.Name, product.Description, product.ID)
	if err != nil {
		return err
	}
	fmt.Println("Product updated successfully!")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
	return nil
}

func DeleteProduct(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		return err
	}

	query := `DELETE FROM products WHERE Id=?;`
	_, err := db.Exec(query, product.ID)
	if err != nil {
		return err
	}
	fmt.Println("Product Deleted successfully!")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
	return nil
}
