package main

import (
	"context"
	"go-microservice-boilerplate/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main(){

	// In this lecture we are going to look at Rest API
	// So what is REST?
	// For most part you will hear simpler definition like rest is transferring JSON over http, 
	// though it not completely wrong.
	// In more descriptive way, Rest stands for REpresentation State Transfer, which is architectural
	// pattern proposed by Roy Fielding. It is most commonly used API pattern to 
	// communicate between services.
	// It has a very specific way of structuring APIs, meaning you will structure 
	// in terms of resources and use HTTP Verbs to create/update/delete/read that 
	// resource(which is usually represented by JSON, but its not must have)
	// mostly you will see JSON being used, as it is lightweight and most of the languages
	// have serializer/deserializer written for it

	// Till now we have seen bookish kind of API examples but going forward we are
	// going to see more professional API, we will build product API for coffee shop
	// we will see about caching, security like additional layer involved
	
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)	
	
	ph := handlers.NewProductHandler(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)	

	s:= &http.Server{
		Addr: ":9090",
		Handler: sm,
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