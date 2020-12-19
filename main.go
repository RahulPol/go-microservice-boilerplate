package main

import (
	"go-microservice-boilerplate/handlers"
	"log"
	"net/http"
	"os"
)

func main(){
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)
	
	// inject the logger into HelloHandler
	hh := handlers.NewHelloHandler(l)

	// In previous lecture we passed nil to http.ListenAndServe, so the 
	// defaultServeMux was used as http.ServeMux. here as well we can do
	// similar thing by calling http.HandleFunc and passing hh and it will add
	// hh.ServeHTTP func as handler for pattern. 

	// But instead we will create a new ServerMux and register our HelloHandler by 
	// calling Handle method. we will see its benifit later
	sm := http.NewServeMux()
	sm.Handle("/", hh)	
	
	// pass newly created serve mux here
	log.Fatal(http.ListenAndServe(":9090", sm))
}