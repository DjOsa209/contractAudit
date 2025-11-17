package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
)

func Invoices(d *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        id := r.URL.Query().Get("id")
        q := "SELECT id,user,amount,note FROM invoices WHERE id = " + id
        rows, _ := d.Query(q)
        defer rows.Close()
        type I struct{ ID int; User string; Amount int; Note string }
        var out []I
        for rows.Next() {
            var x I
            rows.Scan(&x.ID, &x.User, &x.Amount, &x.Note)
            out = append(out, x)
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(out)
    }
}