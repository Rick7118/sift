package database

import "database/sql"

func Tables(db *sql.DB) ([]string, error) {
	rows, err := db.Query(`
		SELECT name
		FROM sqlite_master
		WHERE type = 'table'
		AND name NOT LIKE 'sqlite_%'
		ORDER BY name;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string

	for rows.Next() {
		var name string

		if err := rows.Scan(&name); err != nil {
			return nil, err
		}

		tables = append(tables, name)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tables, nil
}

type Column struct {
	Name       string
	Type       string
	NotNull    bool
	PrimaryKey bool
}

func Schema(db *sql.DB, table string) ([]Column, error) {
	rows, err := db.Query("PRAGMA table_info(" + table + ");")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var columns []Column

	for rows.Next() {
		var (
			cid          int
			name         string
			dataType     string
			notNull      int
			defaultValue any
			primaryKey   int
		)

		if err := rows.Scan(
			&cid,
			&name,
			&dataType,
			&notNull,
			&defaultValue,
			&primaryKey,
		); err != nil {
			return nil, err
		}

		columns = append(columns, Column{
			Name:       name,
			Type:       dataType,
			NotNull:    notNull == 1,
			PrimaryKey: primaryKey == 1,
		})
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}
