package httpapi

import (
	"URezL/Internal/usecases/account"
	"URezL/Internal/usecases/link"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
	"time"
)

const (
	accountIdContextKey = "account_id"
	accountIdUrlPathKey = "account_id"
	shortenLinkUrlPathKey = "shorten_link"
)

type Api struct {
	AccountUseCases account.Interface
	LinkUseCases    link.Interface
}

func CreateApi(a account.Interface, l link.Interface) *Api {
	return &Api{
		AccountUseCases: a,
		LinkUseCases: l,
	}
}

func (a *Api) Router() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/control/cut", a.postLinkCut).Methods(http.MethodPost)
	router.HandleFunc("/control/custom", a.authenticate(a.postCustomLink)).Methods(http.MethodPost)
	router.HandleFunc("/account/login", a.postLogin).Methods(http.MethodPost)
	router.HandleFunc("/account/register", a.postRegister).Methods(http.MethodPost)
	router.HandleFunc("/control/links", a.authenticate(a.getLinks)).Methods(http.MethodGet)
	router.HandleFunc("/control/links/{"+shortenLinkUrlPathKey+"}", a.authenticate(a.deleteLink)).Methods(http.MethodDelete)
	router.HandleFunc("/control/links/{"+shortenLinkUrlPathKey+"}", a.authenticate(a.changeAddress)).Methods(http.MethodPatch)
	router.HandleFunc("/{"+shortenLinkUrlPathKey+"}", a.redirect).Methods(http.MethodGet)


	return router
}

func (a *Api) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lnk, ok := vars[shortenLinkUrlPathKey]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	old, err := a.LinkUseCases.GetLinkByShorten(lnk)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, old, http.StatusSeeOther)
}

type LinkCutRequest struct {
	Link string `json:"link"`
}

// TODO: if authorized get ID
func (a *Api) postLinkCut(w http.ResponseWriter, r *http.Request) {
	l := LinkCutRequest{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	shortenLink, err := a.LinkUseCases.LinkCut(l.Link, nil)
	w.Write([]byte(shortenLink))
	w.WriteHeader(http.StatusCreated)

}

type CustomLinkRequest struct {
	Link string `json:"link"`
	CustomName *string `json:"customName"`
	Lifetime *time.Duration `json:"lifetime"`
}

func (a *Api) postCustomLink(w http.ResponseWriter, r *http.Request) {
	accountId, ok := r.Context().Value(accountIdContextKey).(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	cl := CustomLinkRequest{}
	err := json.NewDecoder(r.Body).Decode(&cl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	customLink, err := a.LinkUseCases.CustomLinkCut(cl.Link, cl.CustomName, cl.Lifetime, accountId)
	w.Write([]byte(customLink))
	w.WriteHeader(http.StatusCreated)
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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/jwt")
	w.Write([]byte(token))
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *Api) postRegister(w http.ResponseWriter, r *http.Request) {
	reg := RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	acc, err := a.AccountUseCases.Register(reg.Username, reg.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(acc.Id))
}

type getLinksResponseModel struct {
	Links     []linkModel `json:"links"`
	LinksNumber int      `json:"links-number"`
}

type linkModel struct {
	Link string `json:"link"`
	CustomName string `json:"customName"`
}

func (a *Api) getLinks(w http.ResponseWriter, r *http.Request) {
	accountId, ok := r.Context().Value(accountIdContextKey).(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	links, err := a.LinkUseCases.GetLinks(accountId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ret := getLinksResponseModel{Links: make([]linkModel, 0, len(links)), LinksNumber: len(links)}
	for _, l := range links {
		ret.Links = append(ret.Links, linkModel{
			Link: l.Link,
			CustomName: l.ShortenLink,
		})
	}
	if err := json.NewEncoder(w).Encode(ret); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (a *Api) deleteLink(w http.ResponseWriter, r *http.Request) {
	accountId, ok := r.Context().Value(accountIdContextKey).(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	vars := mux.Vars(r)
	lnk, ok := vars[shortenLinkUrlPathKey]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err := a.LinkUseCases.DeleteLink(lnk, accountId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))
}

type ChangeAddressRequest struct {
	NewLink string `json:"newLink"`
}

func (a *Api) changeAddress(w http.ResponseWriter, r *http.Request) {
	accountId, ok := r.Context().Value(accountIdContextKey).(string)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ca := ChangeAddressRequest{}
	err := json.NewDecoder(r.Body).Decode(&ca)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	vars := mux.Vars(r)
	lnk, ok := vars[shortenLinkUrlPathKey]
	err = a.LinkUseCases.ChangeAddress(lnk, ca.NewLink, accountId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Address changed"))
}

func (a *Api) authenticate(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearHeader := r.Header.Get("Authorization")
		strArr := strings.Split(bearHeader, " ")
		if len(strArr) != 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token := strArr[1]
		id, err := a.AccountUseCases.Authenticate(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), accountIdContextKey, id)
		handler(w, r.WithContext(ctx))
	}
}


