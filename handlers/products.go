package handlers

import (
	"log"
	"microservices/data"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	// GET
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// POST
	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	// PUT
	if r.Method == http.MethodPut {

		x := regexp.MustCompile(`/([0-9]+)`)
		g := x.FindAllStringSubmatch(r.URL.Path,-1)

		if len(g) != 1{
			http.Error(rw , "Invalid URI",http.StatusBadRequest)
			return 
		}

		if len(g[0]) != 2{
			http.Error(rw , "Invalid URI",http.StatusBadRequest)
			return 
		}

		idString := g[0][1]

		id , err := strconv.Atoi(idString)

		if err != nil {
			http.Error(rw , "Invalid URI",http.StatusBadRequest)
			return 
		}

		p.updateProducts(id , rw  , r)
		return 

	}

	// Catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle POST Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)

}

func (p *Products) updateProducts(id int , rw http.ResponseWriter, r *http.Request ){
	
	p.l.Println("Handle PUT Products")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id , prod)

	if err == data.ErrProductNotFound{
		http.Error(rw , "Product Not Found" , http.StatusNotFound)
		return
	}

	if err != nil{
		http.Error(rw , "Product Not Found" , http.StatusInternalServerError)
		return
	}




}
