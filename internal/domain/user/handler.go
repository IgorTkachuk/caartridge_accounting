package user

import (
	"encoding/json"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	usersURL = "/api/users"
)

type Handler struct {
	UserService Service
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, usersURL, h.GetUsers)
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	logger.Info("Get all users handler")
	users, _ := h.UserService.GetAll(r.Context())

	usersBytes, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(usersBytes)
}
