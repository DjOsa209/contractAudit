package handlers

import (
    "html/template"
    "net/http"
)

func Render() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        t := template.Must(template.New("x").Parse("<div>" + r.URL.Query().Get("txt") + "</div>"))
        t.Execute(w, nil)
    }
}