package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	port := ":8080"
	log.Print("Starting server on ", port)

	err := http.ListenAndServe(port, mux)
	log.Fatal(err)
}
