package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("starting server on port 8080...")
	http.HandleFunc("GET /{message}", echoHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	message := []byte(r.PathValue("message") + "\n")
	log.Printf("received request: %s", message)
	w.Write(message)
}
