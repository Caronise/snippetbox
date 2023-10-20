package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/Caronise/snippetbox/internal/models"

	// We only use the driver's init() function to run so it can register with
	// the database/sql package. Hence, we use a blank identifier.
	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	// Define command-line fag variables and their default values.
	addr := flag.String("addr", ":8080", "HTTP network address to listen on")
	dsn := flag.String("dsn", "web:Poochies8@/snippetbox?parseTime=true",
		"MySQL data source name")

	flag.Parse()

	// Initialize a new structured logger, which writes to the Stdout and use
	// the specified settings.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		//		Level: slog.LevelDebug,
	}))

	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("Starting server on", "addr", *addr)

	// Call app.routes() to get the servemux containing our routes.
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

// OpenDB wraps sql.Open() and returns a sql.DB connection pool for given DSN.
func openDB(dsn string) (*sql.DB, error) {
	// Doesn't create a connection, just initializes the pool.
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	// To verify it's set up correctly, we create a connection with db.Ping()
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
