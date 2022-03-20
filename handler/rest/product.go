package rest

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
)

func (a *RestHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var resp Response
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
