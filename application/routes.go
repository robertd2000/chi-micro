package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/robertd2000/orders-api/handler"
)

func loadRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	orderhandler := &handler.Order{}

	router.Get("/", orderhandler.List)
	router.Post("/", orderhandler.Create)
	router.Get("/{id}", orderhandler.GetById)
	router.Put("/{id}", orderhandler.Update)
	router.Delete("/{id}", orderhandler.Delete)
}
