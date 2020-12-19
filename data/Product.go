package data

import (
	"encoding/json"
	"io"
	"time"
)


type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreateOn    string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error{
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products{
	return productList
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