package usecases

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrPasswordNotCompare = errors.New("password not compare")
	ErrTokenNotValid      = errors.New("token not valid")
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrNotValid           = errors.New("not valid")
)
