package main

import (
	"fmt"
	"io/ioutil"
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
	h1 := func(rw http.ResponseWriter, r *http.Request){
		// the handler function comes with two paramters
		// http.Request -> represent http request. It has properties like Method, URL, Body
		// http.ResponseWriter -> its an interface to construct http response

		// Note that http.ResponseWriter is an interface, so there must be a
		// concrete implementation of this interface.
		// And it is response struct which you will find in http package
		// just follow the trail of http.ListenAndServe(), it has a Server
		// Server has a connection(conn), conn has readRequest and initializes
		// response, and this response implements methods of http.ResponseWriter
		// which are Header(), Write() and WriteHeader()

		// try running curl -X POST -d "Rahul" localhost:9090
		// this creates post request for / path with data "Rahul"
		// to read data received from body
		d, err := ioutil.ReadAll(r.Body)

		if err != nil{
			// This is base method to reply to a host
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Oops"))

			// however http also provides an Error method that will wrap
			// your response writer, message and an error code
			http.Error(rw,"Oops", http.StatusBadRequest)
			
			return
		}

		log.Printf("Data %s\n",d)

		// cool now we can read the data from our API
		// lets try to write data back to our client
		fmt.Fprintf(rw, "Hello %s",d)
		
		log.Println("Hello World")
	}

	// As explained earlier HandleFunc registers handler function for given pattern in the DefaultServeMux
	// so, in simple words we are saying map h1 for all paths 	
	http.HandleFunc("/",h1) 

	
	h2 := func(http.ResponseWriter, *http.Request){
		log.Println("Good Bye")
	}

	// This is just another example of HandleFunc with different pattern(path)
	// Notice that when you request server with goodbye path it won't execute
	// both Hello World and Good Bye 
	http.HandleFunc("/goodbye", h2)
	
	log.Fatal(http.ListenAndServe(":9090", nil))
}