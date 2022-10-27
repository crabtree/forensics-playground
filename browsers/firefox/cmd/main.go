package main

import (
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

	rows, err := db.Query("SELECT prefix || host FROM moz_origins;")
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
