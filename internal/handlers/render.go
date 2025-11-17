package handlers

import (
    "html/template"
    "net/http"
    ttpl "text/template"
)

func Render() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        t := template.Must(template.New("x").Parse("<div>" + r.URL.Query().Get("txt") + "</div>"))
        t.Execute(w, nil)
    }
}

func Render2() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tpl := r.URL.Query().Get("tpl")
        t := ttpl.Must(ttpl.New("y").Parse(tpl))
        t.Execute(w, r.URL.Query())
    }
}