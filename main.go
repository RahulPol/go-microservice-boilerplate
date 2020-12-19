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

	// adding another handler
	gh := handlers.NewGoodByeHandler(l)
	sm.Handle("/goodbye", gh)

	// pass newly created serve mux here
	// log.Fatal(http.ListenAndServe(":9090", sm))

	// with http.ListenAndServe(":9090", sm) we are able to server request over http
	// but still we are not at the point where we want to be
	// since currently we have everything of http server as default like read/write/idle timeout, max header byte
	// we need to tune this parameters so that we don't run into denial of service attack, 
	// so we create our own server with fine tune
	s:= &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	// since we need to cater to multiple request at the same time
	// listen on server parallely
	go func(){
		err := s.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
	}()

	// The following code will make sure that if
	// there is an interrupt to server, it will be 
	// closed gracefully.
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown",sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	
}