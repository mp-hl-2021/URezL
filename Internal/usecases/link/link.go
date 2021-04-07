package link

import "time"

type Link struct {
	Link string
	ShortenLink string
	Lifetime time.Duration
	UserId *int
}

type Interface interface {
	LinkCut(link string) (string, error)
	CustomLinkCut(link string, customName *string, lifetime *time.Duration) (string, error)
	DeleteLink(link string) ()
	ChangeAddress(link string) ()
	GetLinks(userId string) ([]string, error)
}

type UseCases struct{}

func (UseCases) LinkCut(link string) (string, error) {

	panic("not implemented method")
}

func (UseCases) CustomLinkCut(link string, customName *string, lifetime *time.Duration) (string, error) {
	panic("not implemented method")
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