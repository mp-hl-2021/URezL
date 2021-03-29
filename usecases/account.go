package usecases

type Account struct {
	Id string
}

type AccountUseCasesInterface interface {
	Register(login, password string) (Account, error)
	Login(login, password string) (string, error)
	Authenticate(token string) (string, error)
	Logout() ()
}

type AccountUseCases struct {
}

func (AccountUseCases) Register(login string, password string) (Account, error) {
	panic("not implemented method")
}

func (AccountUseCases) Login(login string, password string) (string, error) {
	panic("not implemented method")
}

func (AccountUseCases) Authenticate(token string) (string, error) {
	panic("not implemented method")
}

func (AccountUseCases) Logout() () {
	panic("not implemented method")
}

