package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// we would like to clean our code and not write spaghetti code
// This essentially here is refactoring of the "/" path handler

type HelloHandler struct{
	// keeping logger here so that it can be injected at run time
	// now you can provide any source for your logger it can be
	// a file, a database or std output
	l *log.Logger
}

// NewHelloHandler := initialize hello handler 
func NewHelloHandler(l *log.Logger) *HelloHandler{
	// this is dependency injection
	return &HelloHandler{l}
}

// as explained in previous lecture, any type that implements ServeHttp()
// is http handler. Thus to make HelloHandler handle the http request 
// implement this mehtod
func (h *HelloHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
		if err != nil{			
			http.Error(rw,"Oops", http.StatusBadRequest)			
			return
		}
	fmt.Fprintf(rw, "Hello %s",d)		
}