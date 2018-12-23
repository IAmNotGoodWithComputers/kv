package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func deleteDb(key string) {
	stmt := `DELETE FROM _kv WHERE kv_key = $1`
	_, err := kvDb.Exec(stmt, key)
	if err != nil {
		fmt.Println(fmt.Sprintf("can not delete key %s: %s", key, err.Error()))
		os.Exit(ERR_CANNOT_DELETE_KEY)
	}
}

func searchDb(value string) {
	stmt := `SELECT kv_key FROM _kv WHERE kv_val MATCH $1`
	rows, err := kvDb.Query(stmt, value)
	if err != nil {
		fmt.Println(fmt.Sprintf("can not search for %s: %s", value, err.Error()))
		os.Exit(ERR_CANNOT_READ_VALUE)
	}

	for rows.Next() {
		var key string
		if err = rows.Scan(&key); err != nil {
			fmt.Println(fmt.Sprintf("failed to read key:  %s", err.Error()))
			os.Exit(ERR_CANNOT_READ_VALUE)
		}
		fmt.Println(key)
	}
}

func readKeys() {
	stmt := `SELECT kv_key FROM _kv ORDER BY kv_key ASC`
	rows, err := kvDb.Query(stmt)
	defer rows.Close()
	if err != nil {
		fmt.Println(fmt.Sprintf("can not read keys: %s", err.Error()))
		os.Exit(ERR_CANNOT_READ_KEY)
	}
	for rows.Next() {
		var key string
		if err = rows.Scan(&key); err != nil {
			fmt.Println(fmt.Sprintf("failed to read key:  %s", err.Error()))
			os.Exit(ERR_CANNOT_READ_KEY)
		}
		fmt.Println(key)
	}
}

func saveFromKv(key string, value string) {
	stmt := `REPLACE INTO _kv (kv_key, kv_val) VALUES ($1, $2)`
	_, err := kvDb.Exec(stmt, key, value)
	if err != nil {
		fmt.Println(fmt.Sprintf("can not save key %s: ", err.Error()))
		os.Exit(ERR_CANNOT_SAVE_VALUE)
	}
}

func readKey(key string) {
	stmt := `SELECT kv_val FROM _kv WHERE kv_key = $1`
	res, err := kvDb.Query(stmt, key)
	defer res.Close()
	if err != nil {
		fmt.Println(fmt.Sprintf("failed to read key %s: ", err.Error()))
		os.Exit(ERR_CANNOT_READ_KEY)
	}
	if res.Next() {
		var value string
		if err = res.Scan(&value); err != nil {
			fmt.Println(fmt.Sprintf("failed to read key %s: ", err.Error()))
			os.Exit(ERR_CANNOT_READ_KEY)
		}
		fmt.Println(value)
	}
}

func saveFromStdin(key string) {
	reader := bufio.NewReader(os.Stdin)
	value := ""

	for {
		line, _, err := reader.ReadLine()
		value += string(line) + "\n"
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(fmt.Sprintf("failed to read from stdin: %s", err.Error()))
			os.Exit(ERR_CANNOT_SAVE_VALUE)
		}
	}
	saveFromKv(key, value)
}

func flushDb() {
	_, err := kvDb.Exec("DELETE FROM _kv")
	if err != nil {
		fmt.Println(fmt.Sprintf("can not flush db: %s", err.Error()))
		os.Exit(ERR_CANNOT_FLUSH_DB)
	}
}
