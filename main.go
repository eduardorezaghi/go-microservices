package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// This function will handle all incoming HTTP requests to the root URL.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Root endpoint was called")

		d, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "Request body is required", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Hello, %s!", d)
		http.ResponseWriter.Header(w).Set("Content-Type", "application/json")
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Sayonara, world!"))
		log.Println("Goodbye endpoint was called")
	})

	print("Server is running on port 9090\n")
	// Now, Go will start a web server and listen on port 9090.
	http.ListenAndServe(":9090", nil)
}
