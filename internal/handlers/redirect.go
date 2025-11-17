package handlers

import (
    "net/http"
)

func Redirect() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        to := r.URL.Query().Get("to")
        http.Redirect(w, r, to, http.StatusFound)
    }
}