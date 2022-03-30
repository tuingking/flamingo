package auth

import (
	"errors"
	"time"
)

const Issuer = "Flamingo Authority"

type JwtToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

// JwtClaims implement jwt.Claim interface
type JwtClaims struct {
	// standard claims
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`

	// custom claims
	AccountID   string `json:"aid"`
	IsSuperuser bool   `json:"isu"`
}

func (r JwtClaims) Valid() error {
	now := time.Now().UTC().Unix()

	if now >= r.ExpiresAt {
		return errors.New("token expired")
	}

	if r.Issuer != Issuer {
		return errors.New("invalid issuer")
	}

	return nil
}
