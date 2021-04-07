package main

import (
	"URezL/Internal/interface/httpapi"
	"URezL/Internal/interface/memory/accountrepo"
	"URezL/Internal/service/token"
	"URezL/Internal/usecases/account"
	"flag"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	privateKeyPath := flag.String("privateKey", "app.rsa", "file path")
	publicKeyPath := flag.String("publicKey", "app.rsa.pub", "file path")
	flag.Parse()

	privateKeyBytes, err := ioutil.ReadFile(*privateKeyPath)
	publicKeyBytes, err := ioutil.ReadFile(*publicKeyPath)
	a, err := token.NewKeysRSA(privateKeyBytes, publicKeyBytes, 100*time.Minute)
	if err != nil {
		panic(err)
	}

	accountUseCases := &account.UseCases{
		AccountStorage: accountrepo.NewMemory(),
		Auth:           a,
	}

	service := httpapi.CreateApi(accountUseCases)

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
