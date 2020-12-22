package data

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)


type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreateOn    string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product


func (p *Product) FromJSON(r io.Reader) error{
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Product) Validate() error{
	// for now we are keeping validator in the same struct
	// but for bigger projects its safe to keep a global validator
	// with common configuration 
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return  validate.Struct(p)
}

// This is a custom validator at field level
func validateSKU(fl validator.FieldLevel) bool{
	// sku must have pattern like abc-kkls-lks
	re := regexp.MustCompile(`[a-z]{3}-[a-z]{4}-[a-z]{3}`)
	matches := re.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1{
		return false
	}
	
	return true
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