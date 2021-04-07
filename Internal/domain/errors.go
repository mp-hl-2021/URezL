package domain

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrAlreadyExist = errors.New("already exist")
)
