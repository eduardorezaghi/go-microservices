package main

import (
	"log"
	"net/http"
	"os"

	"github.com/eduardo/Projetos/go-microservices/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	serverMux := http.NewServeMux()
	serverMux.Handle("/", hh)
	serverMux.Handle("/goodbye", gh)

	print("Server is running on port 9090\n")
	// Now, Go will start a web server and listen on port 9090.
	http.ListenAndServe(":9090", serverMux)
}
