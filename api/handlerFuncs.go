package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LinkCutRequest struct {
	Link string `json:"link"`
	Token *string `json:"token"`
}

func PostLinkCut(w http.ResponseWriter, r *http.Request) {
	l := LinkCutRequest{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(l.Link)
	w.Write([]byte("Generated link\n"))
}

type CustomLinkRequest struct {
	Link string `json:"link"`
	CustomName *string `json:"customName"`
	Token string `json:"token"`
	Lifetime *int `json:"lifetime"`
}

// Needs authorizations
func PostCustomLink(w http.ResponseWriter, r *http.Request) {
	cl := CustomLinkRequest{}
	err := json.NewDecoder(r.Body).Decode(&cl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(cl.Link, cl.CustomName)
	w.Write([]byte("Generated custom link\n"))
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func PostLogin(w http.ResponseWriter, r *http.Request) {
	l := LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(l.Username, l.Password)
	w.Write([]byte("Successfully logged in\n"))
}


type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func PostRegister(w http.ResponseWriter, r *http.Request) {
	reg := RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(reg.Username, reg.Password)
	w.Write([]byte("Successfully registered\n"))
}

type GetLinksRequest struct {
	Token string `json:"token"`
}

// Needs authorization
func GetLinks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Links list"))
}

type DeleteLinkRequest struct {
	ShortenLink string `json:"shortenLink"`
	Token string `json:"token"`
}
// Needs authorization
func DeleteLink(w http.ResponseWriter, r *http.Request) {
	dl := DeleteLinkRequest{}
	err := json.NewDecoder(r.Body).Decode(&dl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(dl.ShortenLink)
	w.Write([]byte("Link deleted\n"))
}

type ChangeAddressRequest struct {
	OldCustomLink string `json:"oldCustomLink"`
	NewLink string `json:"newLink"`
	Token string `json:"token"`
}
// Needs authorization
func ChangeAddress(w http.ResponseWriter, r *http.Request) {
	ca := ChangeAddressRequest{}
	err := json.NewDecoder(r.Body).Decode(&ca)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(ca.OldCustomLink, ca.NewLink)
	w.Write([]byte("Link changed\n"))
}


