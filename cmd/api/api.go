package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/VoltealProductions/TheAzureArcchives/middleware"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

type Middleware func(next http.Handler) http.HandlerFunc

func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /hello/{userName}", func(w http.ResponseWriter, r *http.Request) {
		userName := r.PathValue("userName")
		fmt.Println("Hello,", userName, "!")
	})

	v1 := http.NewServeMux()
	v1.Handle("/api/v1", http.StripPrefix("/api/v1", router))

	middlewareChain := MiddlewareChain(
		middleware.RequestLoggerMiddleware,
		middleware.AuthMiddleware,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}

	log.Println("The azure archives is listening on", s.addr)
	return server.ListenAndServe()
}

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next.ServeHTTP
	}
}
