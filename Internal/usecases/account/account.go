package account

import (
	"URezL/Internal/domain/account"
	"URezL/Internal/service/token"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
	"unicode"
)

var (
	ErrInvalidLoginString    = errors.New("login string contains invalid character")
	ErrInvalidPasswordString = errors.New("password string contains invalid character")
	ErrTooShortString        = errors.New("too short string")
	ErrTooLongString         = errors.New("too long string")
	ErrAllCharactersAreUpper = errors.New("must have low characters")
	ErrAllCharactersAreLower = errors.New("must have upper characters")
)

const (
	minLoginLength    = 3
	maxLoginLength    = 100
	minPasswordLength = 5
	maxPasswordLength = 100
)

type Account struct {
	Id string
}

type Interface interface {
	Register(login, password string) (Account, error)
	Login(login, password string) (string, error)
	Authenticate(token string) (string, error)
	Logout() ()
	RegisterLogger(f func(s1, s2 string) (Account, error)) func(s1, s2 string) (Account, error)
	LoginLogger(f func(s1, s2 string) (string, error)) func(s1, s2 string) (string, error)
	AuthenticateLogger(f func(s1 string) (string, error)) func(s1 string) (string, error)
}

type UseCases struct {
	AccountStorage account.Interface
	Auth           token.Interface
}

func (a *UseCases) Register(login string, password string) (Account, error) {
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

	acc, err := a.AccountStorage.CreateAccount(account.Credentials{
		Login:    login,
		Password: string(hashedPassword),
	})
	if err != nil {
		return Account{}, err
	}
	return Account{Id: acc.Id}, nil
}

func (a *UseCases) RegisterLogger(f func(s1, s2 string) (Account, error)) func(s1, s2 string) (Account, error) {
	return func(s1, s2 string) (Account, error){
		start := time.Now()
		acc, ok := f(s1, s2)
		fmt.Printf("method: Register; duration: %v; status: %s\n", time.Since(start), ok)
		return acc, ok
	}
}

func (a *UseCases) Login(login string, password string) (string, error) {
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

func (a *UseCases) LoginLogger(f func(s1, s2 string) (string, error)) func(s1, s2 string) (string, error) {
	return func(s1, s2 string) (string, error){
		start := time.Now()
		s, ok := f(s1, s2)
		fmt.Printf("method: Login; duration: %v; status: %s\n", time.Since(start), ok)
		return s, ok
	}
}

func (a *UseCases) Authenticate(token string) (string, error) {
	return a.Auth.UserIdByToken(token)
}

func (a *UseCases) AuthenticateLogger(f func(s1 string) (string, error)) func(s1 string) (string, error) {
	return func(s1 string) (string, error){
		start := time.Now()
		s, ok := f(s1)
		fmt.Printf("method: Authenticate; duration: %v; status: %s\n", time.Since(start), ok)
		return s, ok
	}
}

func (a *UseCases) Logout() () {
	panic("not implemented method")
}

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

func validatePassword(password string) error {
	chars := 0
	for _, r := range password {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !unicode.IsSpace(r) && !unicode.IsPunct(r){
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
	if strings.ToUpper(password) == password {
		return ErrAllCharactersAreUpper
	}
	if strings.ToLower(password) == password {
		return ErrAllCharactersAreLower
	}
	return nil
}
