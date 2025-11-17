package handlers

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

func TokenNone() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// UnsafeAllowNoneSignatureType 是只读变量，不能直接赋值
		// 这里我们只需要使用 jwt.SigningMethodNone 即可
		claims := jwt.MapClaims{"sub": r.URL.Query().Get("sub")}
		t := jwt.NewWithClaims(jwt.SigningMethodNone, claims)
		s, _ := t.SignedString(jwt.UnsafeAllowNoneSignatureType)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{"token": s})
	}
}
