package main

import (
	"URezL/accountstorage"
	"URezL/auth"
	"URezL/usecases"
	"flag"
	"io/ioutil"
	"net/http"
	"time"

	"URezL/api"
)

func main() {

	privateKeyPath := flag.String("privateKey", "app.rsa", "file path")
	publicKeyPath := flag.String("publicKey", "app.rsa.pub", "file path")
	flag.Parse()

	privateKeyBytes, err := ioutil.ReadFile(*privateKeyPath)
	publicKeyBytes, err := ioutil.ReadFile(*publicKeyPath)
	a, err := auth.NewKeysRSA(privateKeyBytes, publicKeyBytes, 100*time.Minute)
	if err != nil {
		panic(err)
	}

	accountUseCases := &usecases.AccountUseCases{
		AccountStorage: accountstorage.NewMemory(),
		Auth: a,
	}

	service := api.CreateApi(accountUseCases)

	server := http.Server{
		Addr:         "localhost:8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,

		Handler: service.Router(),
	}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
