package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	d, _ := ioutil.ReadAll(r.Body)

	fmt.Fprintf(rw,"Goodbye %s" , d)
}
