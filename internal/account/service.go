package account

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/infra/logger"
)

type Service interface {
	Authenticate(ctx context.Context, username, password string) (Account, error)
	CreateAccount(ctx context.Context, v CreateAccountRequest) (Account, error)
}

type service struct {
	config  Config
	logger  logger.Logger
	account Repository
}

func NewService(config Config, logger logger.Logger, acc Repository) Service {
	return &service{
		config:  config,
		logger:  logger,
		account: acc,
	}
}

func (s *service) Authenticate(ctx context.Context, username, password string) (Account, error) {
	acc, err := s.account.FindByUsername(ctx, username)
	if err != nil {
		return acc, errors.Wrap(err, "find by username")
	}

	match := acc.ComparePassword(password)
	if !match {
		return acc, errors.New("password not match")
	}

	return acc, nil
}

func (s *service) CreateAccount(ctx context.Context, v CreateAccountRequest) (acc Account, err error) {
	hashPassword, err := HashPassword(v.Password)
	if err != nil {
		return acc, errors.Wrap(err, "hash password")
	}

	acc = Account{
		ID:              uuid.New().String(),
		Username:        v.Username,
		password:        hashPassword,
		Name:            v.Name,
		Email:           v.Username,
		Phone:           v.Phone,
		IsEmailVerified: false,
		IsActive:        true,
		IsSuperuser:     v.IsSuperUser,
		Groups:          []Group{},
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		LastLogin:       sql.NullTime{Valid: true, Time: time.Now()},
	}

	err = s.account.CreateAccount(ctx, acc)
	if err != nil {
		return acc, errors.Wrap(err, "create account")
	}

	s.logger.Infof("account created for username: %s", acc.Username)

	// TODO: send email verification

	return acc, nil
}
