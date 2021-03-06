package link

import (
	"time"
)

type Link struct {
	Link        string
	ShortenLink string
	Lifetime    *time.Duration
	AccountId   *string
	Working		bool
}

type Interface interface {
	AddLink(link Link) (Link, error)
	GetLinkByShorten(link string) (Link, error)
	GetLinksByAccountId(accountId string) ([]Link, error)
	CheckLinkExists(link string) bool
	DeleteLink(link string, accountId string) error
	ChangeLink(oldLink string, newLink string, accountId string) error
	SetNotWorking(link string) error
}