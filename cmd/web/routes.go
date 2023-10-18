package main

import "net/http"

// routes returns a servemux containing our application routes.
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	// Create a file server to serve files out of ./ui/static directory.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Register the file server as the handler for all URL paths that start with
	// "/static/". Then strip "/static" prefix.
	// otherwise it will try to access: ./ui/static/static/somefile.css
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Register the other application routes.
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
