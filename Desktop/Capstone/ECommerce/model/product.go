package model

type Product struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
	Price       int    `json:"price,omitempty"`
}
