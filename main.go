package main

import (
	"net/http"
	"time"

	"URezL/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/{shorten_link}", api.Redirect).Methods(http.MethodGet)
	router.HandleFunc("/control/cut", api.PostLinkCut).Methods(http.MethodPost)
	router.HandleFunc("/control/custom", api.PostCustomLink).Methods(http.MethodPost)
	router.HandleFunc("/account/login", api.PostLogin).Methods(http.MethodPost)
	router.HandleFunc("/account/register", api.PostRegister).Methods(http.MethodPost)
	router.HandleFunc("/control/links", api.GetLinks).Methods(http.MethodGet)
	router.HandleFunc("/control/links/{shorten_link}", api.DeleteLink).Methods(http.MethodDelete)
	router.HandleFunc("/control/links/{shorten_link}", api.ChangeAddress).Methods(http.MethodPatch)

	server := http.Server{
		Addr:         "localhost:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
