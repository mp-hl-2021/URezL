package main

import (
	"net/http"
	"time"

	"URezL/api"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/links/cut", api.PostLinkCut).Methods(http.MethodPost)
	router.HandleFunc("/links/custom", api.PostCustomLink).Methods(http.MethodPost)
	router.HandleFunc("/login", api.PostLogin).Methods(http.MethodPost)
	router.HandleFunc("/register", api.PostRegister).Methods(http.MethodPost)
	router.HandleFunc("/links", api.GetLinks).Methods(http.MethodGet)
	router.HandleFunc("/links", api.DeleteLink).Methods(http.MethodDelete)
	router.HandleFunc("/links", api.ChangeAddress).Methods(http.MethodPatch)

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
