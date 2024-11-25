package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type GoodbyeRoute struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *GoodbyeRoute {
	return &GoodbyeRoute{l}
}

func (g *GoodbyeRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye endpoint was called")

	d, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Request body is required", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Goodbye, %s!", d)
	http.ResponseWriter.Header(w).Set("Content-Type", "application-json")
}
