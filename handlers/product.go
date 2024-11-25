package handlers

import (
	"log"
	"net/http"

	"github.com/eduardo/Projetos/go-microservices/data"
)

func NewProducts(
	l *log.Logger,
) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Handle the request
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}