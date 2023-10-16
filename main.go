package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path matches "/". if it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the
	// handler would keep executing.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox"))
}

// Add snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert to an integer. If it can't be converted, or value is less than 1,
	// we return a 404 page not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Fprintf() function to interpolatethe id value with our
	// response and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snipper with ID %d...", id)
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// If not POST, send a 405 status code and "Method not allowed" response body
	if r.Method != http.MethodPost {
		// use the Header().Set() method to indicate what request methods are
		// supported for this URL.
		w.Header().Set("Allow", http.MethodPost)
		// http.Error() calls w.WriteHeader() and w.Write() behind the scenes.
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux, then
	// register each handler to the correponding URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("starting server on:8080")

	// http.ListenAndServe starts a new web server and takes in two parameters:
	// the TCP netwrok address to listen on, and the servemux we just created.
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
