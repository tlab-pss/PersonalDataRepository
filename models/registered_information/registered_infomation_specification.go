package registered_information

import (
	"errors"
	"net/mail"
)

func ValidateMail(m string) error {
	if _, err := mail.ParseAddress(m); err != nil {
		return errors.New("incorrect format")
	}
	return nil
}
