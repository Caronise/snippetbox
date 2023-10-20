package models

import (
	"database/sql"
	"errors"
	"time"
)

// Snippet holds the data for an individual snippet.
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel wraps around a sql.DB connection pool.
type SnippetModel struct {
	DB *sql.DB
}

// Insert will insert a new snippet into the database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	query := `INSERT INTO snippets (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// Use the Exec() method on the embedded db connection pool.
	result, err := m.DB.Exec(query, title, content, expires)
	if err != nil {
		return 0, err
	}

	// result provides two methods: LastInsertId() and RowsAffectd()
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get returns a specific snippet based on its id.
func (m *SnippetModel) Get(id int) (Snippet, error) {
	// Initalize a snippet struct
	var snippet Snippet

	query := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Use the QueryRow() method, passing in the untrusted id variable as the
	// value for the placeholder parameter.
	row := m.DB.QueryRow(query, id)

	// Use row.Scan() to copy the values from each field in sql.Row
	err := row.Scan(&snippet.ID,
		&snippet.Title,
		&snippet.Content,
		&snippet.Created,
		&snippet.Expires,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return snippet, nil
}

// Latest returns the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
