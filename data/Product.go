package data

import (
	"encoding/json"
	"errors"
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


func (p *Product) FromJSON(r io.Reader) error{
	d := json.NewDecoder(r)
	return d.Decode(p)
}

// Notice in this function as well dependency injection is done
// so we can write to any type that implements io.Writer
func (p *Products) ToJSON(w io.Writer) error{
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func UpdateProduct(id int, p *Product) error{
	pos :=  findProduct(id)
	if pos == -1{
		return errors.New("The product with given id does not exist")
	}

	productList[pos] = p
	return nil
}

func findProduct(id int) int{
	pos := -1
	for i,p := range productList{
		if p.ID == id{			
			pos = i
			break
		}
	}

	return pos
}

func AddProduct(p *Product){
	p.ID = getNextProductID()
	p.CreateOn = time.Now().UTC().String()
	p.UpdatedOn = time.Now().UTC().String()
	productList = append(productList, p)
}

func getNextProductID() int{
	if productList == nil{
		return 1
	}

	return productList[len(productList)-1].ID +1
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