package handlers

import (
    "crypto/tls"
    "io"
    "net/http"
)

func Proxy() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        url := r.URL.Query().Get("url")
        c := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
        resp, err := c.Get(url)
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadGateway)
            return
        }
        defer resp.Body.Close()
        io.Copy(w, resp.Body)
    }
}