package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Link struct {
	Link string `json:"link"`
}

func PostLinkCut(w http.ResponseWriter, r *http.Request) {
	l := Link{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(l.Link)
	w.Write([]byte("Generated link\n"))
}

type CustomLink struct {
	Link string `json:"link"`
	CustomName string `json:"customName"`
}
func PostCustomLink(w http.ResponseWriter, r *http.Request) {
	cl := CustomLink{}
	err := json.NewDecoder(r.Body).Decode(&cl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(cl.Link, cl.CustomName)
	w.Write([]byte("Generated custom link\n"))
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func PostLogin(w http.ResponseWriter, r *http.Request) {
	l := Login{}
	err := json.NewDecoder(r.Body).Decode(&l)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(l.Username, l.Password)
	w.Write([]byte("Successfully logged in\n"))
}


type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
func PostRegister(w http.ResponseWriter, r *http.Request) {
	reg := Register{}
	err := json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(reg.Username, reg.Password)
	w.Write([]byte("Successfully registered\n"))
}


// Needs authorizations
func GetLinks(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Links list"))
}

type DeleteLink struct {
	ShortenLink string `json:"shortenLink"`
}
// Needs authorizations
func PostDeleteLink(w http.ResponseWriter, r *http.Request) {
	dl := DeleteLink{}
	err := json.NewDecoder(r.Body).Decode(&dl)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(dl.ShortenLink)
	w.Write([]byte("Link deleted\n"))
}

type ChangeAddress struct {
	OldCustomLink string `json:"oldCustomLink"`
	NewLink string `json:"newLink"`
}
// Needs authorizations
func PostChangeAddress(w http.ResponseWriter, r *http.Request) {
	ca := ChangeAddress{}
	err := json.NewDecoder(r.Body).Decode(&ca)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	fmt.Println(ca.OldCustomLink, ca.NewLink)
	w.Write([]byte("Link changed\n"))
}


