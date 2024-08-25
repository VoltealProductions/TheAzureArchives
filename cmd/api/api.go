package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/VoltealProductions/TheAzureArcchives/middleware"
	"github.com/VoltealProductions/TheAzureArcchives/services/user"
	"github.com/go-chi/chi/v5"
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
	router := chi.NewRouter()

	apiRouter := chi.NewRouter()
	router.Mount("/api/v1", apiRouter)

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(apiRouter)

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
