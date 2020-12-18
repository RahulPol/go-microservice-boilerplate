package main

import (
	"log"
	"net/http"
)

func main(){
		
	
	// Explanation
	// http.ListenAndServe(":9090", nil)
	// this single line is enough to create server,
	// this line translates, go is going to create server that 
	// will listen on port 9090 and is http Handler 
	// since second parameter is nil, it will be using DefaultServeMux
	// which is nothing but instance of ServerMux with default value
	// so it is same as 
	// http.ListenAndServe(":9090", http.DefaultServeMux)

	// ServerMux is an http Handler since it implements ServeHttp() of http.Handler interface
	// this is very important statement as you will see the HandleFunc function later
	// which is mapping of path and handler, which again can be any function which has 
	// same defintion as ServeHttp() and the HandleFunc registers the given path(pattern)
	// in DefaultServerMux

	// In other words, ServeMux is responsible for redirecting a path, so if you provide list
	// of path and function pair(basically handlers) ServeMux will decide which
	// function gets executed

	// you can initialize ServeMux to specific values and use it as well


	// Steps:
	// 1. run the following code 
	// 2. run curl localhost:9090
		// output of above will be 404, as we haven't provided mapping for any
		// path 	
	// 3: run curl localhost:9090
		// now you will get Hello World with time stamp at server
		// in fact for any path you will get the same output
	h1 := func(http.ResponseWriter, *http.Request){
		log.Println("Hello World")
	}

	// As explained earlier HandleFunc registers handler function for given pattern in the DefaultServeMux
	// so, in simple words we are saying map h1 for all paths 	
	http.HandleFunc("/",h1)

	
	log.Fatal(http.ListenAndServe(":9090", nil))
}