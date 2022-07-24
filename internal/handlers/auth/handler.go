package auth

import (
	"encoding/json"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/user"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	authURL   = "/api/auth"
	signupURL = "/api/signup"
)

type Handler struct {
	Logger      logging.Logger
	UserService user.Service
	JWTHelper   jwt.Helper
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodPost, signupURL, apperror.Middleware(h.Signup))
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()
	var dto user.CreateUserDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return apperror.BadRequestError("failed to decode data")
	}

	// TODO после реализации сервиса дописать обработчик
	panic("Implement me")

}
