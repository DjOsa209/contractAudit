package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func Init() *sql.DB {
    d, _ := sql.Open("sqlite3", "vuln.db")
    d.Exec("CREATE TABLE IF NOT EXISTS users(id INTEGER PRIMARY KEY, name TEXT, password TEXT)")
    d.Exec("INSERT INTO users(name, password) VALUES('admin','21232f297a57a5a743894a0e4a801fc3')")
    d.Exec("INSERT INTO users(name, password) VALUES('alice','e1faffb3e614e6c2fba74296962386b7')")
    return d
}

func UnsafeFindUsersByName(d *sql.DB, name string) (*sql.Rows, error) {
    q := "SELECT id, name, password FROM users WHERE name = '" + name + "'"
    return d.Query(q)
}