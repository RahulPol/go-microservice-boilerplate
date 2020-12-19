package handlers

import (
	"go-microservice-boilerplate/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	if r.Method == http.MethodGet{
		p.getProducts(rw,r)
		return
	}	

	if r.Method == http.MethodPost{
		// try running
		// curl -v  -d "{\"name\":\"tea\", \"description\":\"nice cup of tea\", \"sku\":\"ack083\"}" -H "Content-Type: application/json" -X POST http://localhost:9090
		p.addProduct(rw,r)
		return
	}

	if r.Method == http.MethodPut{
		// try running
		// curl -v  -d "{\"name\":\"tea1\", \"description\":\"nice cup of tea\", \"sku\":\"ack083\"}" -H "Content-Type: application/json" -X PUT http://localhost:9090/1
		p.updateProduct(rw, r)
		return
	}

	// catch all
	// try running curl localhost:9090 -X POST -v to reach this code
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// in order to make Rest style API instead of using logic into ServeHTTP 
// lets create separate method for HTTP verbs, so we can provide support
// for HTTP Get/Put/Post/Delete
func (p *ProductHandler) getProducts(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle GET request")
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


func (p *ProductHandler) addProduct(rw http.ResponseWriter, r *http.Request){
	p.l.Println("Handle POST request")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil{
		p.l.Println("Unmarshalling error", err)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	data.AddProduct(prod)		
}
	
func (p *ProductHandler) updateProduct(rw http.ResponseWriter, r *http.Request){
	regex := regexp.MustCompile("/([0-9]+)")
	g := regex.FindAllStringSubmatch(r.URL.Path, -1)		

	if len(g) != 1{
		http.Error(rw,"Invalid parameter: id", http.StatusInternalServerError)	
		return
	}

	if len(g[0]) != 2{
		http.Error(rw,"Invalid parameter: id", http.StatusInternalServerError)	
		return
	}

	idString := g[0][1]
	id,_ := strconv.Atoi(idString)

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil{
		p.l.Println("Unmarshalling error", err)
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	data.UpdateProduct(id, prod)
}