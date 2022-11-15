package errors

import "errors"

var (
	ErrEmailValidation = errors.New("email is not valid")
	ErrDateValidation  = errors.New("date is not valid")
	ErrRoomValidation  = errors.New("room type is not valid")
)
