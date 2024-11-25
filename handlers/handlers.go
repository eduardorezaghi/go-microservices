package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type HelloRoute struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *HelloRoute {
	return &HelloRoute{l}
}

func (h *HelloRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    h.l.Println("Root endpoint was called")

    d, err := io.ReadAll(r.Body)

    if err != nil {
        http.Error(w, "Request body is required", http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "Hello, %s!", d)
    http.ResponseWriter.Header(w).Set("Content-Type", "application/json")
}
