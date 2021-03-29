package usecases

import "time"

type LinkUseCasesInterface interface {
	LinkCut(link string) (string, error)
	CustomLinkCut(link string, customName *string, lifetime *time.Duration) (string, error)
	DeleteLink(link string) ()
	ChangeAddress(link string) ()
	GetLinks(userId string) ([]string, error)
}

type LinkUseCases struct{}

func (LinkUseCases) LinkCut(link string) (string, error) {
	panic("not implemented method")
}

func (LinkUseCases) CustomLinkCut(link string, customName *string, lifetime *time.Duration) (string, error) {
	panic("not implemented method")
}

func (LinkUseCases) DeleteLink(link string) () {
	panic("not implemented method")
}

func (LinkUseCases) ChangeAddress(link string) () {
	panic("not implemented method")
}

func (LinkUseCases) GetLinks(userId string) ([]string, error) {
	panic("not implemented method")
}