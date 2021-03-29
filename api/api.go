package api

import (
	"URezL/usecases"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Api struct {
	AccountUseCases usecases.AccountUseCasesInterface
	LinkUseCases    usecases.LinkUseCasesInterface
}

func CreateApi(a usecases.AccountUseCasesInterface, l usecases.LinkUseCasesInterface) *Api {
	return &Api{
		AccountUseCases: a,
		LinkUseCases:    l,
	}
}

func (a *Api) Router() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/{shorten_link}", a.redirect).Methods(http.MethodGet)
	router.HandleFunc("/control/cut", a.postLinkCut).Methods(http.MethodPost)
	router.HandleFunc("/control/custom", a.postCustomLink).Methods(http.MethodPost)
	router.HandleFunc("/account/login", a.postLogin).Methods(http.MethodPost)
	router.HandleFunc("/account/register", a.postRegister).Methods(http.MethodPost)
	router.HandleFunc("/control/links", a.getLinks).Methods(http.MethodGet)
	router.HandleFunc("/control/links/{shorten_link}", a.deleteLink).Methods(http.MethodDelete)
	router.HandleFunc("/control/links/{shorten_link}", a.changeAddress).Methods(http.MethodPatch)

	return router
}

func (a *Api) redirect(w http.ResponseWriter, r *http.Request) {
	// todo: change on redirection
	w.Write([]byte("Redirect to real link\n"))
}

type Link struct {
	Link string
	ShortenLink string
	Lifetime time.Duration
	UserId *int
}

type LinkCutRequest struct {
	Link string `json:"link"`
	Token *string `json:"token"`
}

func (a* Api) postLinkCut(w http.ResponseWriter, r *http.Request) {
	l := LinkCutRequest{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	shortenLink, err := a.LinkUseCases.LinkCut(l.Link)
	fmt.Println(shortenLink)
	w.Write([]byte("Generated link\n"))
}

type CustomLinkRequest struct {
	Link string `json:"link"`
	CustomName *string `json:"customName"`
	Token string `json:"token"`
	Lifetime *time.Duration `json:"lifetime"`
}

// Needs authorizations
func (a *Api) postCustomLink(w http.ResponseWriter, r *http.Request) {
	cl := CustomLinkRequest{}
	err := json.NewDecoder(r.Body).Decode(&cl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	customLink, err := a.LinkUseCases.CustomLinkCut(cl.Link, cl.CustomName, cl.Lifetime)
	fmt.Println(customLink)
	w.Write([]byte("Generated custom link\n"))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func (a *Api) postLogin(w http.ResponseWriter, r *http.Request) {
	l := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	token, err := a.AccountUseCases.Login(l.Username, l.Password)
	fmt.Println(token)
	w.Write([]byte("Successfully logged in\n"))
}


type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func (a * Api) postRegister(w http.ResponseWriter, r *http.Request) {
	reg := RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	token, err := a.AccountUseCases.Register(reg.Username, reg.Password)
	fmt.Println(token)
	w.Write([]byte("Successfully registered\n"))
}

type GetLinksRequest struct {
	Token string `json:"token"`
}

// Needs authorization
func (a *Api) getLinks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Links list"))
}

type DeleteLinkRequest struct {
	Token string `json:"token"`
}

// Needs authorization
func (a *Api) deleteLink(w http.ResponseWriter, r *http.Request) {
	dl := DeleteLinkRequest{}
	err := json.NewDecoder(r.Body).Decode(&dl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	// todo: get link from address
	a.LinkUseCases.DeleteLink("Todo: link from address")
	w.Write([]byte("Link deleted\n"))
}

type ChangeAddressRequest struct {
	NewLink string `json:"newLink"`
	Token string `json:"token"`
}

// Needs authorization
func (a *Api) changeAddress(w http.ResponseWriter, r *http.Request) {
	ca := ChangeAddressRequest{}
	err := json.NewDecoder(r.Body).Decode(&ca)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	// todo: get link from address
	a.LinkUseCases.ChangeAddress("Todo: link from address")
	w.Write([]byte("Link changed\n"))
}


