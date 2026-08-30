package database

import (
	"strings"
)

type QueryResult struct {
	Columns      []string
	Rows         [][]any
	RowsAffected int64
}

func (db *Database) Execute(query string) (*QueryResult, error) {
	query = strings.TrimSpace(query)

	if query == "" {
		return nil, nil
	}

	firstWord := strings.ToUpper(strings.Fields(query)[0])

	switch firstWord {
	case "SELECT", "PRAGMA", "WITH":
		return db.executeQuery(query)

	default:
		return db.executeStatement(query)
	}
}

func (db *Database) executeQuery(query string) (*QueryResult, error) {
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := &QueryResult{
		Columns: columns,
		Rows:    make([][]any, 0),
	}

	for rows.Next() {
		values := make([]any, len(columns))
		pointers := make([]any, len(columns))

		for i := range values {
			pointers[i] = &values[i]
		}

		if err := rows.Scan(pointers...); err != nil {
			return nil, err
		}

		result.Rows = append(result.Rows, values)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (db *Database) executeStatement(query string) (*QueryResult, error) {
	result, err := db.DB.Exec(query)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &QueryResult{
		RowsAffected: rowsAffected,
	}, nil
}
