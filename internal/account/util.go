package account

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return password, errors.Wrap(err, "bcrypt failed generate from password")
	}

	return string(bytes), nil
}
