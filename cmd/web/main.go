package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define command-line fag variables and their default values.
	addr := flag.String("addr", ":8080", "HTTP network address to listen on")

	flag.Parse()

	mux := http.NewServeMux()

	// Create a file server to serve files out of ./ui/static directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". Then strip "/static" prefix.
	// otherwise it will try to access: ./ui/static/static/somefile.css
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Register the other application routes.
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on ", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
