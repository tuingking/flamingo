package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	requestID := r.Context().Value(middleware.RequestIDKey).(string)
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

// GetAllProducts
// @Summary get all products
// @Description get all products
// @Tags Product
// @Accept json
// @Produce json
// @Param paramName query string true "my param"
// @Success 200 {object} Response{} "Success Response"
// @Failure 400 "Bad Request"
// @Failure 500 "InternalServerError"
// @Router /products/bulk [post]
func (h *RestHandler) BulkCreate(w http.ResponseWriter, r *http.Request) {
	var resp Response
	defer resp.Render(w, r)

	type request struct {
		Filename string `json:"filename"`
	}
	var req request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.SetError(errors.New("decode request body"), http.StatusBadRequest)
		return
	}

	if req.Filename == "" {
		resp.SetError(errors.New("filename required"), http.StatusBadRequest)
		return
	}
	h.logger.Info("filename: ", req.Filename)

	err = h.product.BulkCreate(r.Context(), req.Filename)
	if err != nil {
		resp.SetError(errors.Wrap(err, "bulk create"), http.StatusInternalServerError)
		return
	}

	resp.Data = nil
}
