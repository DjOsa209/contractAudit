package handlers

import "net/http"

func CorsEcho() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("ok"))
    }
}