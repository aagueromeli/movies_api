package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HealthRouter() chi.Router {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return router
}
