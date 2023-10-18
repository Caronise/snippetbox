package main

import (
	"net/http"
)

// serverError writes a log entry at Error level (including the request method
// and URI as attributes), then sends a generic 500 Internal Server Error
// response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	//	trace  = string(debug.Stack())
	)

	app.logger.Error(err.Error(),
		"method", method,
		"uri", uri,
	// 	"trace", trace,
	)
	http.Error(w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

// clientError sends a response including the specific status code and
// corresponding description to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// notFound is a convenience wrapper around clientError which sends a 404
// Not Found response to the user.
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
