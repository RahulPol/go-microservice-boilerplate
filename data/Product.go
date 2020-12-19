package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreateOn    string
	UpdatedOn   string
	DeletedOn   string
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc232",
		CreateOn:    time.Now().UTC().String(),
		UpdatedOn:    time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fdk32",
		CreateOn:    time.Now().UTC().String(),
		UpdatedOn:    time.Now().UTC().String(),
	},
}