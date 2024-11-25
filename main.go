package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	// create a new server
	s := &http.Server{
		Addr:         ":9090", 		 // configure the bind address
		Handler:      serverMux, // set the default handler
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		ReadTimeout:  1 * time.Second,  // max time to read request from the client
		WriteTimeout: 1 * time.Second, // max time to write response to the client
	}

	// start the server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// consumes the signal from the channel
	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	s.Shutdown(tc)
}
