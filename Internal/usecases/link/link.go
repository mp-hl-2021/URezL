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
	shortenLink := generateShortenLink(lnk)
	l, err := a.LinkStorage.AddLink(link.Link{
			Link: lnk,
			ShortenLink: shortenLink,
			Lifetime: 100 * time.Second,
			UserId: accountId,
		})
	if err != nil {
		return "", err
	}
	return l.ShortenLink, nil
}

func (a *UseCases) CustomLinkCut(lnk string, customName *string, lifetime *time.Duration, accountId string) (string, error) {
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
		UserId: &accountId,
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

func generateShortenLink(link string) (hash string) {
	//TODO: add hash generation
	 return "aaa"
}