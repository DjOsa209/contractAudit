package handlers

import (
    "io"
    "net/http"
    "os"
)

func FileRead() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        p := r.URL.Query().Get("path")
        http.ServeFile(w, r, "./uploads/"+p)
    }
}

func FileUpload() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        f, _ := os.Create("./uploads/" + name)
        defer f.Close()
        io.Copy(f, r.Body)
        w.WriteHeader(http.StatusCreated)
    }
}