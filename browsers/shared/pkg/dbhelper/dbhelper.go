package dbhelper

import (
	"database/sql"
	"fmt"
	"os"
)

func LoadDB(dbFilePath string) (*sql.DB, error) {
	_, err := os.Stat(dbFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("Database file does not exists under path %s\n", dbFilePath)
		}
		return nil, fmt.Errorf("An error while reading database file: %s", err)
	}

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, fmt.Errorf("An error while opening the database file: %s", err)
	}
	return db, nil
}
