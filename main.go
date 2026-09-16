package main

import (
	"fmt"
	"learn-go/application"
	"net/http"
)

func main() {
	router := application.LoadRoutes()
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("could not start server: %v", err)
	}
}
