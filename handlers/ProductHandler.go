package handlers

import (
	"go-microservice-boilerplate/data"
	"log"
	"net/http"
)

type ProductHandler struct{
	l *log.Logger
}

// NewProductHandler := initialize product handler 
func NewProductHandler(l *log.Logger) *ProductHandler{
	// this is dependency injection
	return &ProductHandler{l}
}

func (p *ProductHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	products := data.GetProducts()

	// One of the method to marshal json
	// d, err := json.Marshal(products)
	// if err != nil{
	// 	http.Error(rw,"Unable to marshal json", http.StatusInternalServerError)
	// }

	// rw.Write(d)

	// other and better approach, better because it wont store the marshalled json
	// rather it will directly write it to rw 
	err := products.ToJSON(rw)
	if err != nil{
		http.Error(rw,"Unable to marshal json", http.StatusInternalServerError)
	}	
}