package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
	"github.com/tuingking/flamingo/infra/contextkey"
	"github.com/tuingking/flamingo/internal/auth"
	"github.com/tuingking/flamingo/internal/product"
)

// GetAllProducts
// @Summary get all products
// @Description get all products
// @Tags Product
// @Accept json
// @Produce json
// @Param paramName query string true "my param"
// @Success 200 {object} Response{data=[]product.Product} "Success Response"
// @Failure 400 "Bad Request"
// @Failure 500 "InternalServerError"
// @Router /products [get]
func (h *RestHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var resp Response
	defer resp.Render(w, r)

	//*** How to get chi requestID
	requestID, ok := r.Context().Value(middleware.RequestIDKey).(string)
	if !ok {
		resp.SetError(errors.New("unable to get request id"), http.StatusInternalServerError)
		return
	}
	h.logger.Info("requestID: ", requestID)

	//** Get Session
	fmt.Printf("[DEBUG] r.Context().Value(contextkey.Identity): %+v\n", r.Context().Value(contextkey.Identity))
	identity, ok := r.Context().Value(contextkey.Identity).(auth.JwtClaims)
	if !ok {
		h.logger.Warn("unable to get identity")
	}
	fmt.Printf("[DEBUG] identity: %+v\n", identity)

	products, err := h.product.GetAllProducts(r.Context())
	if err != nil {
		resp.SetError(errors.Wrap(err, "get all products"), http.StatusInternalServerError)
		return
	}
	var respData []product.Product = products

	resp.Data = respData
}

// Seed Product
// @Summary auto generate random product and populate it to database
// @Description auto generate random product and populate it to database
// @Tags Product
// @Accept json
// @Produce json
// @Param n path integer true "number product are going to be generated"
// @Success 200 {object} Response{} "Success Response"
// @Failure 400 "Bad Request"
// @Failure 500 "InternalServerError"
// @Router /products/seed/{n} [post]
func (h *RestHandler) SeedProduct(w http.ResponseWriter, r *http.Request) {
	var resp Response
	defer resp.Render(w, r)

	n, err := strconv.Atoi(chi.URLParam(r, "n"))
	if err != nil {
		resp.SetError(errors.Wrap(err, "n should be a number"), http.StatusInternalServerError)
		return
	}

	err = h.product.Seed(r.Context(), n)
	if err != nil {
		resp.SetError(errors.Wrap(err, "[Seed] error seeding product data"), http.StatusInternalServerError)
		return
	}

	resp.Data = nil
}
