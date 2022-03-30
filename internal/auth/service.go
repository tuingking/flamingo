package auth

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/internal/account"
)

type Service interface {
	IssueJwtToken(ctx context.Context, acc account.Account) (JwtToken, error)
	VerifyAccessToken(accessToken string) (JwtClaims, error)
}

type service struct {
	config Config
	logger logger.Logger
}

type Config struct {
	PrivateKey string
}

func NewService(config Config, logger logger.Logger) Service {
	return &service{
		config: config,
		logger: logger,
	}
}

func (s *service) IssueJwtToken(ctx context.Context, acc account.Account) (tkn JwtToken, err error) {
	const (
		accessTokenExpiresIn  = 1 * time.Hour
		refreshTokenExpiresIn = 76 * time.Hour
	)

	// Access Token
	accessClaims := JwtClaims{
		AccountID:   acc.ID,
		IsSuperuser: acc.IsSuperuser,
		Issuer:      Issuer,
		IssuedAt:    time.Now().Unix(),
		ExpiresAt:   time.Now().Add(accessTokenExpiresIn).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	at, err := accessToken.SignedString([]byte(s.config.PrivateKey))
	if err != nil {
		return tkn, errors.Wrap(err, "signing access token")
	}

	// Refresh Token
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expires_in": time.Now().Add(refreshTokenExpiresIn).Unix(),
	})
	rt, err := refreshToken.SignedString([]byte(s.config.PrivateKey))
	if err != nil {
		return tkn, errors.Wrap(err, "signing refresh token")
	}

	tkn = JwtToken{
		AccessToken:  at,
		RefreshToken: rt,
		ExpiresIn:    int64(accessTokenExpiresIn),
	}

	// TODO: save to DB

	return tkn, err
}

func (s *service) VerifyAccessToken(accessToken string) (claims JwtClaims, err error) {
	if accessToken == "" {
		return claims, errors.New("no access token found")
	}

	// Authorization: Bearer {tokenstr}
	authScheme := accessToken[:6]
	tokenstr := accessToken[7:]

	if strings.ToUpper(authScheme) != "BEARER" && tokenstr != "" {
		return claims, errors.Wrap(err, "invalid access token")
	}

	// parse tokenstr
	_, err = jwt.ParseWithClaims(tokenstr, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.config.PrivateKey), nil
	})
	if err != nil {
		return claims, errors.Wrap(err, "parse claims")
	}

	if err = claims.Valid(); err != nil {
		return claims, errors.Wrap(err, "invalid claims")
	}

	return claims, nil
}
