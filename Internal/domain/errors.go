package domain

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrAlreadyExist = errors.New("already exist")
	ErrPermissionDenied = errors.New("permission denied")
)
