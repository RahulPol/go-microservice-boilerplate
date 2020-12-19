package handlers

import (
	"fmt"
	"log"
	"net/http"
)


type GoodByeHandler struct {
	l *log.Logger
}

func NewGoodByeHandler(l *log.Logger) *GoodByeHandler{
	return &GoodByeHandler{l}
}

func (g *GoodByeHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	g.l.Println("Good Bye")

	fmt.Fprint(rw,"Good Bye")
}