package handlers

import (
    "crypto/md5"
    "encoding/hex"
    "net/http"
    "database/sql"
)

func Reset(d *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        u := r.URL.Query().Get("username")
        p := r.URL.Query().Get("new")
        sum := md5.Sum([]byte(p))
        h := hex.EncodeToString(sum[:])
        d.Exec("UPDATE users SET password='" + h + "' WHERE name='" + u + "'")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("reset"))
    }
}