package link

import (
	"URezL/Internal/domain/link"
	"math/rand"
	"time"
)

const basicLifetime = 100 * time.Minute
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"


type Link struct {
	Link string
	ShortenLink string
}

type Interface interface {
	GetLinkByShorten(link string) (string, error)
	LinkCut(link string, accountId *string) (string, error)
	CustomLinkCut(link string, customName *string, lifetime *time.Duration, accountId string) (string, error)
	DeleteLink(link string, accountId string) error
	ChangeAddress(oldLink string, newLink string, accountId string) error
	GetLinks(accountId string) ([]Link, error)
}

type UseCases struct{
	LinkStorage link.Interface
}

func (a *UseCases) GetLinkByShorten(lnk string) (string, error) {
	l, err := a.LinkStorage.GetLinkByShorten(lnk)
	if err != nil {
		return "", err
	}
	return l.Link, nil
}

func (a *UseCases) LinkCut(lnk string, accountId *string) (string, error) {
	shortenLink := generateShortenLink(a)
	dur := 100 * time.Second
	l, err := a.LinkStorage.AddLink(link.Link{
			Link:        lnk,
			ShortenLink: shortenLink,
			Lifetime:    &dur,
			AccountId:   accountId,
		})
	if err != nil {
		return "", err
	}
	return l.ShortenLink, nil
}

func (a *UseCases) CustomLinkCut(lnk string, customName *string, lifetime *time.Duration, accountId string) (string, error) {
	if customName == nil {
		tmp := generateShortenLink(a)
		customName = &tmp
	}
	if lifetime == nil {
		tmp := basicLifetime
		lifetime = &tmp
	}
	l, err := a.LinkStorage.AddLink(link.Link{
		Link:        lnk,
		ShortenLink: *customName,
		Lifetime:    lifetime,
		AccountId:   &accountId,
	})
	if err != nil {
		return "", err
	}
	return l.ShortenLink, nil
}

func (a *UseCases) DeleteLink(link string, accountId string) error {
	err := a.LinkStorage.DeleteLink(link, accountId)
	return err
}

func (a *UseCases) ChangeAddress(oldLink string, newLink string, accountId string) error {
	err := a.LinkStorage.ChangeLink(oldLink, newLink, accountId)
	return err
}

func (a *UseCases) GetLinks(accountId string) ([]Link, error) {
	links, err := a.LinkStorage.GetLinksByAccountId(accountId)
	if err != nil {
		return nil, err
	}
	res := make([]Link, 0, len(links))
	for _, l := range links {
		res = append(res, Link{
			Link: l.Link,
			ShortenLink: l.ShortenLink,
		})
	}
	return res, nil
}

func generateShortenLink(a *UseCases) (hash string) {
	curLength := 1
	randString := RandStringBytes(curLength)
	for  a.LinkStorage.CheckLinkExists(randString){
		curLength += 1
		randString = RandStringBytes(curLength)
	}
	return randString
}

func RandStringBytes(n int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}