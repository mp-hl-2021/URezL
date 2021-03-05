package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"URezL/api"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/linkCut", api.PostLinkCut).Methods(http.MethodPost)
	router.HandleFunc("/customLink", api.PostCustomLink).Methods(http.MethodPost)
	router.HandleFunc("/login", api.PostLogin).Methods(http.MethodPost)
	router.HandleFunc("/register", api.PostRegister).Methods(http.MethodPost)
	router.HandleFunc("/links", api.GetLinks).Methods(http.MethodGet)
	router.HandleFunc("/deleteLink", api.PostDeleteLink).Methods(http.MethodPost)
	router.HandleFunc("/changeAddress", api.PostChangeAddress).Methods(http.MethodPost)

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

func getTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	o := struct {
		Hello string `json:"hello"`
	}{
		Hello: "world",
	}
	if err := json.NewEncoder(w).Encode(o); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}