package link

import (
	"time"
)

type Link struct {
	Link string
	ShortenLink string
	Lifetime time.Duration
	UserId *string
}

type Interface interface {
	AddLink(link Link) (Link, error)
	GetLinkByShorten(link string) (Link, error)
	GetLinksByAccountId(accountId string) ([]Link, error)
	CheckLinkExists(link string) (bool, error)
}