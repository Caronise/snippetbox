package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Caronise/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	// Initialize a slice containing the paths to the template files.
	// NOTE: the base template MUST be the first file in the slice!
	//files := []string{
	//	"./ui/html/base.tmpl",
	//	"./ui/html/pages/home.tmpl",
	//	"./ui/html/partials/nav.tmpl",
	//}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v", snippet)
	}

	// Use template.ParseFiles() function to read the files and store them into
	// a set. Pass the files slice as variadic parameters.
	//ts, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, r, err)
	//	return
	//}

	// Use the ExecuteTemplate() method to write the content of the "base"
	// template as the response body. The last paramater to ExecuteTemplate()
	// represents dynamic data that can be passed in, which can be nil.
	//err = ts.ExecuteTemplate(w, "base", nil)
	//if err != nil {
	//	app.serverError(w, r, err)
	//}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// Use SnippetModel's Get() method to retrieve the data based on ID.
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	// Write the snippet data as plaint-text http response body.
	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// dummy vars, delete this later
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n - Kobayashi Issa"
	expires := 7

	// pass data to SnippetModel.Insert(), receiving the ID of the new record.
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
	}

	// Redirect user to relevant page for the new snippet
	snippetPath := fmt.Sprintf("/snippet/view?id=%d", id)
	http.Redirect(w, r, snippetPath, http.StatusSeeOther)
}
