package cartridge_model

import (
	"encoding/json"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/cartridge_model"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	cartridgeModelsUrl = "/api/ctrmodel"
	cartridgeModelUrl  = "/api/ctrmodel/:id"
)

var _ handlers.Handler = &Handler{}

type Handler struct {
	CartridgeModelSvc cartridge_model.Service
}

func (h Handler) Register(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, cartridgeModelsUrl, jwt.Middleware(apperror.Middleware(h.GetCarteidgeModels)))
}

func (h Handler) GetCarteidgeModels(w http.ResponseWriter, r *http.Request) error {
	logger := logging.GetLogger()
	logger.Info("Execute Get all printers handler")

	models, err := h.CartridgeModelSvc.GetAll(r.Context())
	if err != nil {
		return err
	}
	modelsBytes, err := json.Marshal(models)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(modelsBytes)
	return nil
}
