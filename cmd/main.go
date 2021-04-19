package main

import (
	"URezL/Internal/interface/httpapi"
	"URezL/Internal/interface/postgres/accountrepo"
	"URezL/Internal/interface/postgres/linkrepo"
	"URezL/Internal/service/token"
	"URezL/Internal/usecases/account"
	"URezL/Internal/usecases/link"
	_ "github.com/lib/pq"

	"encoding/json"
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

	filename := "cmd/config/config.postgres.json"
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	type Configuration struct {
		ConnectionString string `json:"connectionString"`
	}
	config := Configuration{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	connStr := config.ConnectionString
	accountUseCases := &account.UseCases{
		AccountStorage: accountrepo.New(connStr),
		Auth:           a,
	}

	linkUseCases := &link.UseCases{
		LinkStorage: linkrepo.New(connStr),
	}

	service := httpapi.CreateApi(accountUseCases, linkUseCases)

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
