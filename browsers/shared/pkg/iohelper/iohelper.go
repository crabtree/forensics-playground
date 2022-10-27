package iohelper

import (
	"log"
	"os"
)

func PrintUsageAndExit() {
	log.Fatalf("Usage: %s <path to database file>\n", os.Args[0])
}

func ExitOnError(err error) {
	if err == nil {
		return
	}
	log.Fatalln(err)
}
