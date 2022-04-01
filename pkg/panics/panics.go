package panics

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"

	"github.com/eapache/go-resiliency/breaker"
)

var (
	// circuit breaker
	cb *breaker.Breaker
)

// HTTPRecoveryMiddleware act as middleware that capture panics standard in http handler
func HTTPRecoveryMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		_, _ = httputil.DumpRequest(r, true)
		defer func() {
			if !recoveryBreak() {
				rcv := panicRecover(recover())
				if rcv != nil {
					// log the panic
					fmt.Fprintf(os.Stderr, "Panic: %+v\n", rcv)
					debug.PrintStack()
					// publishError(rcv, request, true)
					http.Error(w, rcv.Error(), http.StatusInternalServerError)
				}
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func recoveryBreak() bool {
	if cb == nil {
		return false
	}

	if err := cb.Run(func() error {
		return nil
	}); err == breaker.ErrBreakerOpen {
		return true
	}
	return false
}

func panicRecover(rc interface{}) error {
	if cb != nil {
		r := cb.Run(func() error {
			return recovery(rc)
		})
		return r
	}
	return recovery(rc)
}

func recovery(r interface{}) error {
	var err error
	if r != nil {
		switch t := r.(type) {
		case string:
			err = errors.New(t)
		case error:
			err = t
		default:
			err = errors.New("Unknown error")
		}
	}
	return err
}
