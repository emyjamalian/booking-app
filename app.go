package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func app() {

	router := chi.NewRouter()
	router.Get("/hello", basicHandler)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	serverErr := server.ListenAndServe()
	if serverErr != nil {
		fmt.Println("failed to listen to the server", serverErr)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
