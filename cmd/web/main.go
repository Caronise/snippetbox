package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// Define command-line fag variables and their default values.
	addr := flag.String("addr", ":8080", "HTTP network address to listen on")

	flag.Parse()

	// Initialize a new structured logger, which writes to the Stdout and use
	// the specified settings.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		//		Level: slog.LevelDebug,
	}))

	app := &application{
		logger: logger,
	}

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

	logger.Info("Starting server on", slog.String("addr", *addr))

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
