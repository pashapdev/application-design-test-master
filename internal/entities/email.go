package entities

import "net/mail"

type Email string

func (e Email) Valid() error {
	_, err := mail.ParseAddress(string(e))
	if err != nil {
		return err
	}
	return nil
}
