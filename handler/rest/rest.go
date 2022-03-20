package rest

import (
	"html/template"

	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/internal/product"
)

type RestHandler struct {
	logger  logger.Logger
	tpl     *template.Template
	product product.Service
}

func NewRestHandler(
	logger logger.Logger,
	tpl *template.Template,
	product product.Service,
) RestHandler {
	return RestHandler{
		logger:  logger,
		tpl:     tpl,
		product: product,
	}
}
