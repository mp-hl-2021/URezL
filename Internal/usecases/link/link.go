package link

import (
	"URezL/Internal/domain/link"
	"time"
)

const basicLifetime = 100 * time.Minute

type Link struct {
	Link string
	ShortenLink string
}

type Interface interface {
	LinkCut(link string) (string, error)
	CustomLinkCut(link string, customName *string, lifetime *time.Duration) (string, error)
	DeleteLink(link string) ()
	ChangeAddress(link string) ()
	GetLinks(userId string) ([]string, error)
}

type UseCases struct{
	LinkStorage link.Interface
}

func (a *UseCases) LinkCut(lnk string, userId *string) (string, error) {
	shortenLink := generateShortenLink(lnk)
	l, err := a.LinkStorage.AddLink(link.Link{
			Link: lnk,
			ShortenLink: shortenLink,
			Lifetime: 100 * time.Second,
			UserId: userId,
		})
	if err != nil {
		return "", err
	}
	return l.ShortenLink, nil
}

func (a *UseCases) CustomLinkCut(lnk string, customName *string, lifetime *time.Duration, userId *string) (string, error) {
	if customName == nil {
		tmp := generateShortenLink(lnk)
		customName = &tmp
	}
	if lifetime == nil {
		tmp := basicLifetime
		lifetime = &tmp
	}
	l, err := a.LinkStorage.AddLink(link.Link{
		Link: lnk,
		ShortenLink: *customName,
		Lifetime: *lifetime,
		UserId: userId,
	})
	if err != nil {
		return "", err
	}
	return l.ShortenLink, nil
}

func (UseCases) DeleteLink(link string) () {
	panic("not implemented method")
}

func (UseCases) ChangeAddress(link string) () {
	panic("not implemented method")
}

func (UseCases) GetLinks(userId string) ([]string, error) {
	panic("not implemented method")
}

func generateShortenLink(link string) (hash string) {
	// TODO: generate only correct link
	panic("not implemented")
}