package server

import (
    "contractAudit/internal/handlers"
    "contractAudit/internal/db"
    "database/sql"
    "net/http"
)

type Server struct {
    mux *http.ServeMux
    db  *sql.DB
}

func New() *Server {
    s := &Server{mux: http.NewServeMux()}
    s.db = db.Init()
    s.routes()
    return s
}

func (s *Server) Start() {
    http.ListenAndServe(":8080", s.cors(s.mux))
}

func (s *Server) routes() {
    s.mux.HandleFunc("/login", handlers.Login())
    s.mux.HandleFunc("/users", handlers.Users(s.db))
    s.mux.HandleFunc("/file", handlers.FileRead())
    s.mux.HandleFunc("/upload", handlers.FileUpload())
    s.mux.HandleFunc("/exec", handlers.Exec())
    s.mux.HandleFunc("/proxy", handlers.Proxy())
    s.mux.HandleFunc("/render", handlers.Render())
    s.mux.HandleFunc("/render2", handlers.Render2())
    s.mux.HandleFunc("/debug", handlers.Debug())
    s.mux.HandleFunc("/redirect", handlers.Redirect())
    s.mux.HandleFunc("/cors", handlers.CorsEcho())
    s.mux.HandleFunc("/ws", handlers.WS())
    s.mux.HandleFunc("/jwtcheck", handlers.JwtCheck())
}

func (s *Server) cors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        w.Header().Set("Access-Control-Allow-Headers", "*")
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        next.ServeHTTP(w, r)
    })
}