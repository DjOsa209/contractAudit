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
    d.Exec("CREATE TABLE IF NOT EXISTS invoices(id INTEGER PRIMARY KEY, user TEXT, amount INTEGER, note TEXT)")
    d.Exec("INSERT INTO invoices(user, amount, note) VALUES('alice', 100, 'Service fee')")
    d.Exec("INSERT INTO invoices(user, amount, note) VALUES('admin', 999, 'VIP charge')")
    d.Exec("CREATE TABLE IF NOT EXISTS contracts(id INTEGER PRIMARY KEY, owner TEXT, title TEXT, content TEXT, status TEXT)")
    d.Exec("INSERT INTO contracts(owner, title, content, status) VALUES('alice','Demo Contract','<b>Welcome</b>','draft')")
    return d
}

func UnsafeFindUsersByName(d *sql.DB, name string) (*sql.Rows, error) {
    q := "SELECT id, name, password FROM users WHERE name = '" + name + "'"
    return d.Query(q)
}