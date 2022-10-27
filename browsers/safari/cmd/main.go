package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/crabtree/forensics-playground/browsers/shared/pkg/dbhelper"
	"github.com/crabtree/forensics-playground/browsers/shared/pkg/iohelper"
)

func main() {
	if len(os.Args) != 2 {
		iohelper.PrintUsageAndExit()
	}

	dbFilePath := os.Args[1]
	db, err := dbhelper.LoadDB(dbFilePath)
	iohelper.ExitOnError(err)
	defer db.Close()

	rows, err := db.Query("SELECT url FROM history_items;")
	iohelper.ExitOnError(err)

	hosts := make(map[string]int)
	for rows.Next() {
		var address string
		rows.Scan(&address)

		u, err := url.Parse(address)
		if err != nil {
			log.Println("Unable to parse string as URL: %s", address)
		}

		if _, exists := hosts[u.Host]; !exists {
			hosts[u.Host] = 0
		}
		hosts[u.Host] += 1
	}

	fmt.Printf("%+v\n", hosts)
}

func printUsageAndExit() {
	log.Fatalf("Usage: %s <path to History.db>\n", os.Args[0])
}

func loadDB(dbFilePath string) (*sql.DB, error) {
	_, err := os.Stat(dbFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("History.db file does not exists under path %s\n", dbFilePath)
		}
		return nil, fmt.Errorf("An error while reading History.db file: %s", err)
	}

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		return nil, fmt.Errorf("An error while opening the database file: %s", err)
	}
	return db, nil
}

func exitOnError(err error) {
	if err == nil {
		return
	}
	log.Fatalln(err)
}
