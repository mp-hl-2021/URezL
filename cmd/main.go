package main

import (
	"URezL/Internal/interface/httpapi"
	"URezL/Internal/interface/postgres/accountrepo"
	"URezL/Internal/interface/postgres/linkrepo"
	"URezL/Internal/service/token"
	"URezL/Internal/usecases/account"
	"URezL/Internal/usecases/link"

	"database/sql"
	_ "github.com/lib/pq"

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

	//TODO: add config file
	connStr := "user=postgres password=123455678 host=localhost dbname=postgres sslmode=disable"
	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	accountUseCases := &account.UseCases{
		AccountStorage: accountrepo.New(conn),
		Auth:           a,
	}

	linkUseCases := &link.UseCases{
		LinkStorage: linkrepo.New(conn),
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
