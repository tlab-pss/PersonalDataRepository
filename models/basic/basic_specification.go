package basic

import (
	"errors"
)

func ValidateName(n string) error {
	if len(n) < 1 || 25 < len(n) {
		return errors.New("name length should be 2 ~ 24")
	}
	return nil
}

func ValidateGender(g int) error {
	// 0: male, 1: female, 2: other
	if g != 0 && g != 1 && g != 2 {
		return errors.New("gender should be 0 ~ 1")
	}
	return nil
}
