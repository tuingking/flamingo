package rest

import (
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Response defines http response for the client
type Response struct {
	Code       int         `json:"code"`
	Data       interface{} `json:"data,omitempty"`
	Error      Error       `json:"error"`
	Message    string      `json:"message"`
	ServerTime int64       `json:"serverTime"`
	Pagination interface{} `json:"pagination,omitempty"`
}

// Error defines the error
type Error struct {
	Status bool   `json:"status" example:"false"` // true if we have error
	Msg    string `json:"msg" example:" "`        // error message
	Code   int    `json:"code" example:"0"`       // application error code for tracing
}

// SetError set the response to return the given error.
// code is http status code, http.StatusInternalServerError is the default value
func (res *Response) SetError(err error, code ...int) {
	logrus.Error(errors.Wrap(err, "ERR"))
	res.ServerTime = time.Now().Unix()

	if len(code) > 0 {
		res.Code = code[0]
	} else {
		res.Code = http.StatusInternalServerError
	}

	if err != nil {
		res.Error = Error{
			Msg:    err.Error(),
			Status: true,
		}
	}

}

// Render writes the http response to the client
func (res *Response) Render(w http.ResponseWriter, r *http.Request) {
	res.ServerTime = time.Now().Unix()

	if res.Code == 0 {
		res.Code = http.StatusOK
	}

	render.Status(r, res.Code)
	render.JSON(w, r, res)
}
