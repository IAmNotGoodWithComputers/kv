package main

import (
	"database/sql"
	"fmt"
	"os"
)

var kvDb *sql.DB

func main() {
	if len(os.Args) == 1 {
		showHelp()
		os.Exit(0)
	} else if len(os.Args) == 2 {
		if os.Args[1] == "--flush" {
			flushDb()
		} else if os.Args[1] == "--keys" {
			readKeys()
		}else {
			stat, _ := os.Stdin.Stat()
			if (stat.Mode() & os.ModeCharDevice) == 0 {
				saveFromStdin(os.Args[1])
			} else {
				readKey(os.Args[1])
			}
		}
	} else if len(os.Args) == 3 && os.Args[1] == "--search" {
		searchDb(os.Args[2])
	} else if len(os.Args) == 3 && os.Args[1] == "--delete" {
		deleteDb(os.Args[2])
	} else if len(os.Args) == 3 {
		saveFromKv(os.Args[1], os.Args[2])
	} else {
		showHelp()
	}
}

func showHelp() {
	fmt.Println(`kv [github.com/tiananmensquare/kv]
a utility to save temporary data to an internal database

usage:

	kv [key]
		reads a key from the internal db
	kv [key] [value]
		saves [value] under the key [key] into the internal db
	kv --flush
		flushes the internal db
	kv --keys
		prints all saved keys in alphabetical order
	kv --delete somekey
		removes the key from the internal database
	kv --search sometext
		reads all keys that match 'sometext'

		note that matching is done on a word basis, to partially match a 
		word, use wildcards (kv --search 'somet*')

additional functions:

	somecommand | kv [key]
		saves stdin under [key] into the internal db

source code:
	https://github.com/tiananmensquare/kv

LICENSE: MIT`)
}
