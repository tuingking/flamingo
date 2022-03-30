package rest

import (
	"html/template"

	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/internal/account"
	"github.com/tuingking/flamingo/internal/auth"
	"github.com/tuingking/flamingo/internal/product"
)

type RestHandler struct {
	logger logger.Logger
	tpl    *template.Template

	// service
	auth    auth.Service
	product product.Service
	account account.Service
}

func NewRestHandler(
	logger logger.Logger,
	tpl *template.Template,

	// service
	auth auth.Service,
	product product.Service,
	account account.Service,
) RestHandler {
	return RestHandler{
		logger:  logger,
		tpl:     tpl,
		auth:    auth,
		product: product,
		account: account,
	}
}
