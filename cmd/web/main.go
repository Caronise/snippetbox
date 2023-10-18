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

	logger.Info("Starting server on", "addr", *addr)

	// Call app.routes() to get the servemux containing our routes.
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
