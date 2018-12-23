### a utility to save temporary data to an internal database

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

notes:

relies on SQLite3 with FTS4. SQLite 3 needs to be installed on your system
and it needs to support FTS4. Currently only works on GNU/Linux systems (and
maybe on macos systems with SQLite installed via brew?)

the internal database is in `$HOME/.kv.db` and can be queries via SQLite

installation:

    go get github.com/tiananmensquare/kv
    go install github.com/tiananmensquare/kv
    
