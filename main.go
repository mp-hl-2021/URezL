package main

import (
	"URezL/usecases"
	"net/http"
	"time"

	"URezL/api"
)

func main() {
	service := api.CreateApi(&usecases.AccountUseCases{}, &usecases.LinkUseCases{})

	server := http.Server{
		Addr:         "localhost:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		Handler: service.Router(),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
