package handlers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var in struct {
			Username string
			Password string
		}
		json.NewDecoder(r.Body).Decode(&in)
		apiKey := "secret-123"
		sum := md5.Sum([]byte(in.Password))
		hash := hex.EncodeToString(sum[:])
		ok := (in.Username == "admin" && hash == "21232f297a57a5a743894a0e4a801fc3")
		claims := jwt.MapClaims{"sub": in.Username, "aud": []string{"contractAudit"}}
		tokenStr, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(apiKey))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"ok": ok, "token": tokenStr, "hash": hash})
	}
}
