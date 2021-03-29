package usecases

import (
	"URezL/accountstorage"
	"URezL/auth"

	"golang.org/x/crypto/bcrypt"

	"errors"
	"unicode"
)

var (
	ErrInvalidLoginString    = errors.New("login string contains invalid character")
	ErrInvalidPasswordString = errors.New("password string contains invalid character")
	ErrTooShortString        = errors.New("too short string")
	ErrTooLongString         = errors.New("too long string")

	ErrInvalidLogin    		 = errors.New("login not found")
	ErrInvalidPassword 		 = errors.New("invalid password")
)

// TODO: change validation constants
const (
	minLoginLength    = 6
	maxLoginLength    = 20
	minPasswordLength = 14
	maxPasswordLength = 48
)

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
	AccountStorage accountstorage.Interface
	Auth           auth.Interface
}

func (a *AccountUseCases) Register(login string, password string) (Account, error) {
	if err := validateLogin(login); err != nil {
		return Account{}, err
	}
	if err := validatePassword(password); err != nil {
		return Account{}, err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return Account{}, err
	}
	acc, err := a.AccountStorage.CreateAccount(accountstorage.Credentials{
		Login:    login,
		Password: string(hashedPassword),
	})
	if err != nil {
		return Account{}, err
	}
	return Account{Id: acc.Id}, nil
}

func (a *AccountUseCases) Login(login string, password string) (string, error) {
	if err := validateLogin(login); err != nil {
		return "", err
	}
	if err := validatePassword(password); err != nil {
		return "", err
	}
	acc, err := a.AccountStorage.GetAccountByLogin(login)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(acc.Credentials.Password), []byte(password)); err != nil {
		return "", err
	}
	token, err := a.Auth.IssueToken(acc.Id)
	if err != nil {
		return "", err
	}
	return token, err
}

func (a *AccountUseCases) Authenticate(token string) (string, error) {
	return a.Auth.UserIdByToken(token)
}

func (a *AccountUseCases) Logout() () {
	panic("not implemented method")
}

// TODO: change validation rules
func validateLogin(login string) error {
	chars := 0
	for _, r := range login {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return ErrInvalidLoginString
		}
		chars++
	}
	if chars < minLoginLength {
		return ErrTooShortString
	}
	if chars > maxLoginLength {
		return ErrTooLongString
	}
	return nil
}

// TODO: change validation rules
func validatePassword(password string) error {
	chars := 0
	for _, r := range password {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) {
			return ErrInvalidPasswordString
		}
		chars++
	}
	if chars < minPasswordLength {
		return ErrTooShortString
	}
	if chars > maxPasswordLength {
		return ErrTooLongString
	}
	return nil
}
