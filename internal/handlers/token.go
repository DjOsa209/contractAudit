package handlers

import (
    "encoding/json"
    "net/http"
    jwt "github.com/dgrijalva/jwt-go"
)

func TokenNone() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        jwt.UnsafeAllowNoneSignatureType = true
        claims := jwt.MapClaims{"sub": r.URL.Query().Get("sub")}
        t := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
        s, _ := t.SignedString(jwt.UnsafeAllowNoneSignatureType)
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]any{"token": s})
    }
}