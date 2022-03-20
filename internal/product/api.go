package product

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/infra/logger"
	"github.com/tuingking/flamingo/internal/apis"
)

type API struct {
	logger  logger.Logger
	product Service
}

func NewAPI(logger logger.Logger, product Service) API {
	return API{
		logger:  logger,
		product: product,
	}
}

func (a *API) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var resp apis.Response
	defer resp.Render(w, r)

	//*** How to get chi requestID
	requestID := r.Context().Value(middleware.RequestIDKey).(string)
	a.logger.Info("requestID: ", requestID)

	products, err := a.product.GetAllProducts(r.Context())
	if err != nil {
		resp.SetError(errors.Wrap(err, "get all products"), http.StatusInternalServerError)
	}

	resp.Data = products
}
