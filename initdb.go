package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/user"
)

func init() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(fmt.Sprintf("kv can not find the currently logged in user: %s", err.Error()))
		os.Exit(ERR_USER_NOT_FOUND)
	}

	dbFile := fmt.Sprintf("%s/.kv.db", usr.HomeDir)
	needsInit := false

	_, fErr := os.Stat(dbFile)
	if fErr != nil && os.IsNotExist(fErr) {
			needsInit = true
		_, fcErr := os.Create(dbFile)
		if fcErr != nil {
			fmt.Println(fmt.Sprintf("can not create db file under %s: %s", dbFile, fcErr.Error()))
			os.Exit(ERR_CANNOT_CREATE_DBFILE)
		}
	} else if fErr != nil && !os.IsNotExist(fErr) {
		fmt.Println(fmt.Sprintf("can not read db file under %s: %s", dbFile, fErr.Error()))
		os.Exit(ERR_DBFILE_UNREADABLE)
	}

	kvDb, err = sql.Open("sqlite3", fmt.Sprintf("file:%s", dbFile))
	if err != nil {
		fmt.Println(fmt.Sprintf("can not open db file for SQL access: %s", err.Error()))
		os.Exit(ERR_CANNOT_OPEN_DBFILE)
	}

	if needsInit {
		q := `CREATE VIRTUAL TABLE _kv USING FTS4 (
			kv_key TEXT PRIMARY KEY,
			kv_val TEXT NOT NULL
		)`

		_, sqlErr := kvDb.Exec(q)
		if sqlErr != nil {
			fmt.Println(fmt.Sprintf("can not init db schema: %s", sqlErr.Error()))
			os.Exit(ERR_CANNOT_INIT_SCHEMA)
		}
	}
}