package user

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
	usersURL = "/api/users"
)

type Handler struct {
	UserService user.Service
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, jwt.Middleware(apperror.Middleware(h.GetUsers)))
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) error {
	logger := logging.GetLogger()
	logger.Info("Get all users handler")
	users, err := h.UserService.GetAll(r.Context())
	if err != nil {
		return err
	}

	usersBytes, err := json.Marshal(users)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usersBytes)

	return nil
}
