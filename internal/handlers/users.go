package handlers

import (
    "contractAudit/internal/db"
    "database/sql"
    "encoding/json"
    "net/http"
)

func Users(d *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        rows, err := db.UnsafeFindUsersByName(d, name)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()
        type U struct{ ID int; Name string; Password string }
        var list []U
        for rows.Next() {
            var u U
            rows.Scan(&u.ID, &u.Name, &u.Password)
            list = append(list, u)
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(list)
    }
}