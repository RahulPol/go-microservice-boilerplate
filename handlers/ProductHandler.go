package handlers

import (
	"context"
	"fmt"
	"go-microservice-boilerplate/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductHandler struct{
	l *log.Logger
}

// NewProductHandler := initialize product handler 
func NewProductHandler(l *log.Logger) *ProductHandler{
	// this is dependency injection
	return &ProductHandler{l}
}

// in order to make Rest style API instead of using logic into ServeHTTP 
// lets create separate method for HTTP verbs, so we can provide support
// for HTTP Get/Put/Post/Delete
func (p *ProductHandler) GetProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle GET request")
	products := data.GetProducts()

	// One of the method to marshal json
	// d, err := json.Marshal(products)
	// if err != nil{
	// 	http.Error(rw,"Unable to marshal json", http.StatusInternalServerError)
	// }
	// rw.Write(d)

	// other and better approach, better because it won't store the marshalled json
	// rather it will directly write it to rw 
	err := products.ToJSON(rw)
	if err != nil{
		http.Error(rw,"Unable to marshal json", http.StatusInternalServerError)
	}	
}


func (p *ProductHandler) AddProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request")

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	data.AddProduct(prod)		
}
	
func (p *ProductHandler) UpdateProduct(rw http.ResponseWriter, r *http.Request){
	// we are using vars here instead of complex logic
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil{
		http.Error(rw, "Unable to convert id", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(*data.Product)
	
	data.UpdateProduct(id, prod)
}

type KeyProduct struct{}

func (p ProductHandler) MiddlewareProductValidation(next http.Handler) http.Handler{
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request){
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil{
			p.l.Println("Unmarshalling error", err)
			http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
			return
		}

		// validate the product
		if err = prod.Validate(); err != nil{
			p.l.Println("[Error] Validating Product", err)
			http.Error(
				rw, 
				fmt.Sprintf("[Error] Validating Product: %s", err), 
				http.StatusBadRequest,
			)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw,req)
	})

}