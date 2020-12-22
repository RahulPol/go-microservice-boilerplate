package main

import (
	"context"
	"go-microservice-boilerplate/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main(){

	// In this lecture we are going to look at validating your API input
	// We are going to use go-playground/validator tool (https://github.com/go-playground/validator)
	// Validation is really important to achieve security of your API
	// If you look at the OWASP's Top 10 vulnerability (https://sucuri.net/guides/owasp-top-10-security-vulnerabilities-2020/)
	// you will understand its very important to sanitize your data input
	// in addition to securing the API, the validator also provides you a nice way to give correct error message back to caller
	// We are going to validate our Product data model in this lecture

	l := log.New(os.Stdout, "product-api ", log.LstdFlags)	
	
	ph := handlers.NewProductHandler(l)

	router := mux.NewRouter()

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)
	
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareProductValidation)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareProductValidation)
	

	s:= &http.Server{
		Addr: ":9090",
		Handler: router,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}
	
	go func(){
		err := s.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown",sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)	
}