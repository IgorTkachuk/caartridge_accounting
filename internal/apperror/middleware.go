package apperror

import (
	"errors"
	"net/http"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h AppHandler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *AppError

		err := h(w, r)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())
				}
				err := err.(*AppError)
				w.WriteHeader(http.StatusBadRequest)
				w.Write(err.Marshal())
				return
			}
			w.WriteHeader(http.StatusTeapot)
			w.Write(systemError(err.Error()).Marshal())
		}
	}
}
