package application

import (
	"learn-go/application/handler"
	"learn-go/application/storage"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func LoadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Route("/Tickets", loadTicketRoutes)

	return router
}

func loadTicketRoutes(router chi.Router) {
	TicketHandler := &handler.Ticket{Repo: &storage.TicketStorage{}}

	router.Post("/", TicketHandler.Create)
	router.Get("/", TicketHandler.List)
	router.Get("/{id}", TicketHandler.GetByID)
	router.Put("/{id}", TicketHandler.UpdateByID)
	router.Delete("/{id}", TicketHandler.DeleteByID)
}
