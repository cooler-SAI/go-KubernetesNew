package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello Gopher from Kubernetes!")
		if err != nil {
			log.Printf("Error writing response: %v", err)
			return
		}
		log.Printf("Received request from %s", r.RemoteAddr)
	})

	port := "8080"
	log.Printf("Server starting on port %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
