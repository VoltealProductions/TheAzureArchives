package server

import (
	"net/http"

	"github.com/VoltealProductions/TheAzureArcchives/utils"
	"github.com/go-chi/chi/v5"
)

func RegisterBaseRoutes(router *chi.Mux) *chi.Mux {
	router.Group(func(r chi.Router) {
		r.HandleFunc("GET /", handleGetIndex)
	})

	return router
}

func handleGetIndex(w http.ResponseWriter, r *http.Request) {
	data := ExampleType{
		Test: "Hello, World!",
		Base: BasePage{
			PageTitle: "Index",
			SiteTitle: "The Azure Archives",
		},
	}

	utils.RenderHtml(w, r, "index", data)
}
