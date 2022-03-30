package rest

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/internal/account"
	"github.com/tuingking/flamingo/internal/auth"
)

// CreateAccount
// @Summary 	Create user account
// @Description Create user account
// @Tags 		Account
// @Accept 		json
// @Produce 	json
// @Param 		request body account.CreateAccountRequest true "payload to create account"
// @Success 	200 {object} Response{data=account.Account} "Success Response"
// @Failure 	400 "Bad Request"
// @Failure 	500 "InternalServerError"
// @Router /accounts [post]
func (h *RestHandler) RegisterNewAccount(w http.ResponseWriter, r *http.Request) {
	var (
		req account.CreateAccountRequest
		res = Response{
			Data: account.Account{},
		}
	)
	defer res.Render(w, r)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res.SetError(errors.Wrap(err, "Unable to unmarshal request body"), http.StatusBadRequest)
		return
	}

	acc, err := h.account.CreateAccount(r.Context(), req)
	if err != nil {
		res.SetError(errors.Wrap(err, "failed to create user account"), http.StatusInternalServerError)
		return
	}

	res.Data = acc
}

// Issue Access Token
// @Summary 	Issue Access Token
// @Description Issue Access Token
// @Tags 		Account
// @Accept 		json
// @Produce 	json
// @Param 		request body account.SignInRequest true "payload to get access token"
// @Success 	200 {object} Response{data=auth.JwtToken} "Success Response"
// @Failure 	400 "Bad Request"
// @Failure 	500 "InternalServerError"
// @Router /auth/token [post]
func (h *RestHandler) IssueAccessToken(w http.ResponseWriter, r *http.Request) {
	var (
		req account.SignInRequest
		res = Response{
			Data: auth.JwtToken{},
		}
	)
	defer res.Render(w, r)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		res.SetError(errors.Wrap(err, "Unable to unmarshal request body"), http.StatusBadRequest)
		return
	}

	acc, err := h.account.Authenticate(r.Context(), req.Username, req.Password)
	if err != nil {
		res.SetError(errors.Wrap(err, "invalid username or password"), http.StatusBadRequest)
		return
	}

	jwtToken, err := h.auth.IssueJwtToken(r.Context(), acc)
	if err != nil {
		res.SetError(errors.Wrap(err, "failed create access token"), http.StatusInternalServerError)
		return
	}

	res.Data = jwtToken
}
