package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	h1 := func(rw http.ResponseWriter, r *http.Request){
		d, err := ioutil.ReadAll(r.Body)
		if err != nil{			
			http.Error(rw,"Oops", http.StatusBadRequest)			
			return
		}
		fmt.Fprintf(rw, "Hello %s",d)		
	}
	
	h2 := func(http.ResponseWriter, *http.Request){
		log.Println("Good Bye")
	}

	http.HandleFunc("/",h1) 
	http.HandleFunc("/goodbye", h2)
	
	log.Fatal(http.ListenAndServe(":9090", nil))
}