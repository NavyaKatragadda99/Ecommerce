package api

import (
	"ECommerce/dataservice"
	"ECommerce/model"
	"database/sql"
	"net/http"
)

func AddLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.AddProduct(db, w, r)
}

func ListLogic(db *sql.DB) ([]model.Product, error) {
	return dataservice.ListProduct(db)
}

func UpdateLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.UpdateProduct(db, w, r)
}
func DeleteLogic(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	return dataservice.DeleteProduct(db, w, r)
}
