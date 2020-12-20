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

	// In this lecture we are going to look at Gorilla Mux  which is a third party multiplexer. 
	// till now we were using default server mux to handle api request, which is good and it
	// satisfy the need but in this case we end up creating a lot of boilerplate code for eg
	// like we have to handle parameter parsing from query string, handle http verb type, parsing 
	// the data etc

	// The gorilla web tookkit is package that contains a lot of stuff other than just mux but we
	// are focused on gorilla/mux package for this lecture
	// detailed info on gorilla mux at https://github.com/gorilla/mux

	// however to give you a gist, we can
	// 1. Register route paths to handlers (similar to servermux), like serveMux it also has handlefunc and handle
	// 2. Register paths can have variables
		// eg. r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler) // here category and id are variables with id having regex 
	// 3. Restricted routes to domain or subdomain
		// eg. r.Host("www.example.com") // Only matches if domain is "www.example.com".
		// eg. r.Host("{subdomain:[a-z]+}.example.com") // Matches a dynamic subdomain.
	// 4. There are several other matchers that can be added
	// 4.1. Match path prefix
		// r.PathPrefix("/products/")
	// 4.2. or HTTP methods
		// r.Methods("GET", "POST")
	// 4.3. or URL schemes
		// r.Schemes("https")
	// 4.4. or header values
		// r.Headers("X-Requested-With", "XMLHttpRequest")
	// 4.5. or to use a custom matcher function
		// r.MatcherFunc(func(r *http.Request, rm *RouteMatch) bool {
		// 	return r.ProtoMajor == 0
		// })
	// and finally, it is possible to combine several matchers in a single route
		// r.HandleFunc("/products", ProductsHandler).
		// Host("www.example.com").
		// Methods("GET").
		// Schemes("http")

	// Routes are tested in the order they were added to the router. If two routes match, the first one wins
	// Setting the same matching conditions again and again can be boring, so we have a way to group several 
	// routes that share the same requirements. We call it "subrouting"
	// eg let's say we have several URLs that should only match when the host is www.example.com. 
	// Create a route for that host and get a "subrouter" from it
	// r := mux.NewRouter()
	// s := r.Host("www.example.com").Subrouter()

	// Then register routes in the subrouter
	// 	s.HandleFunc("/products/", ProductsHandler)
	// s.HandleFunc("/products/{key}", ProductHandler)
	// s.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)


	// Another important feature of gorilla mux is it supports Middleware
	// Middlewares are (typically) small pieces of code which take one request, do something with it, and pass 
	// it down to another middleware or the final handler. Some common use cases for middleware are request logging, 
	// header manipulation, or ResponseWriter hijacking.

	// Mux supports the addition of middlewares to a Router, which are executed in the order they are added if a match 
	// is found, including its subrouters.

	// Mux middlewares are defined using the de facto standard type:
	// type MiddlewareFunc func(http.Handler) http.Handler

	// a simple example
	// func loggingMiddleware(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		// Do stuff here
	// 		log.Println(r.RequestURI)
	// 		// Call the next handler, which can be another middleware in the chain, or the final handler.
	// 		next.ServeHTTP(w, r)
	// 	})
	// }

	// r := mux.NewRouter()
	// r.HandleFunc("/", handler)
	// r.Use(loggingMiddleware)


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