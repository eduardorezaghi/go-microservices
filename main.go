package main

import "net/http"

func main() {
	// This function will handle all incoming HTTP requests to the root URL.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Now, Go will start a web server and listen on port 9090.
	http.ListenAndServe(":9090", nil)
}
