package server

import (
	"contractAudit/internal/db"
	"contractAudit/internal/handlers"
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
	s.mux.HandleFunc("/invoices", handlers.Invoices(s.db))
	s.mux.HandleFunc("/debug", handlers.Debug())
	s.mux.HandleFunc("/redirect", handlers.Redirect())
	s.mux.HandleFunc("/cors", handlers.CorsEcho())
	s.mux.HandleFunc("/ws", handlers.WS())
	s.mux.HandleFunc("/reset", handlers.Reset(s.db))
	s.mux.HandleFunc("/token-none", handlers.TokenNone())
	s.mux.HandleFunc("/contracts/create", handlers.ContractsCreate(s.db))
	s.mux.HandleFunc("/contracts/get", handlers.ContractsGet(s.db))
	s.mux.HandleFunc("/contracts/search", handlers.ContractsSearch(s.db))
	s.mux.HandleFunc("/contracts/sign", handlers.ContractsSign(s.db))
	s.mux.HandleFunc("/contracts/assign", handlers.ContractsAssign(s.db))
	s.mux.HandleFunc("/contracts/upload", handlers.ContractsUpload())
	s.mux.HandleFunc("/contracts/review", handlers.ContractsReview())
	s.mux.HandleFunc("/contracts/approve", handlers.ContractsApprove(s.db))
	s.mux.HandleFunc("/contracts/download", handlers.ContractsDownload())
	s.mux.HandleFunc("/contracts/webhook", handlers.ContractsWebhook(s.db))
	s.mux.HandleFunc("/contracts/delete", handlers.ContractsDelete(s.db))
	s.mux.HandleFunc("/contracts/signurl", handlers.ContractsSignURL())
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
