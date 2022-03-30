package account

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	ID              string       `json:"id"`
	Username        string       `json:"username"`
	password        string       `json:"-"`
	Name            string       `json:"name"`
	Email           string       `json:"email"`
	Phone           string       `json:"phone"`
	IsEmailVerified bool         `json:"is_email_verified"`
	IsActive        bool         `json:"is_active"`
	IsSuperuser     bool         `json:"is_superuser"`
	Groups          []Group      `json:"group"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	LastLogin       sql.NullTime `json:"last_login"`
}

func NewAccount(username, password string) Account {
	return Account{
		ID:       uuid.New().String(),
		Username: username,
		password: password,
	}
}

func (r *Account) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(r.password), []byte(password))
	return err == nil
}

type CreateAccountRequest struct {
	Username    string `json:"username" default:"youremail@gmail.com"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	IsSuperUser bool   `json:"is_superuser" default:"false"`
}

type UpdatePasswordRequest struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ForgotPasswordRequest struct {
	Username string `json:"username"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
