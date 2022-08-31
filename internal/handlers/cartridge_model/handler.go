package cartridge_model

import (
	"encoding/json"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/cartridge_model"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
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
	r.HandlerFunc(http.MethodGet, cartridgeModelUrl, jwt.Middleware(apperror.Middleware(h.GetCartridgeModelById)))
	r.HandlerFunc(http.MethodPost, cartridgeModelsUrl, jwt.Middleware(apperror.Middleware(h.CreateCartridgeModel)))
	r.HandlerFunc(http.MethodPatch, cartridgeModelsUrl, jwt.Middleware(apperror.Middleware(h.UpdateCartridgeModel)))
	r.HandlerFunc(http.MethodDelete, cartridgeModelUrl, jwt.Middleware(apperror.Middleware(h.DeleteCartridgeModel)))
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

func (h Handler) GetCartridgeModelById(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	modelId, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	ctrModel, err := h.CartridgeModelSvc.GetById(r.Context(), modelId)
	if err != nil {
		return err
	}

	ctrModelBytes, err := json.Marshal(ctrModel)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ctrModelBytes)

	return nil
}

func (h Handler) CreateCartridgeModel(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var newCtrModel cartridge_model.CreateCartridgeModelDTO
	err := json.NewDecoder(r.Body).Decode(&newCtrModel)
	if err != nil {
		return apperror.BadRequestError("Failed to decode new cartridge model data")
	}

	newCtrModelId, err := h.CartridgeModelSvc.Create(r.Context(), newCtrModel)
	if err != nil {
		return err
	}

	w.Header().Set("Location", fmt.Sprintf("%s/%s", cartridgeModelsUrl, newCtrModelId))
	return nil
}

func (h Handler) UpdateCartridgeModel(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var updCtrModelInfo cartridge_model.UpdateCartridgeModelDTO
	if err := json.NewDecoder(r.Body).Decode(&updCtrModelInfo); err != nil {
		return apperror.BadRequestError("Failed to decode info for update cartridge model")
	}

	err := h.CartridgeModelSvc.Update(r.Context(), updCtrModelInfo)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h Handler) DeleteCartridgeModel(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	ctrModelId, _ := strconv.Atoi(params.ByName("id"))

	err := h.CartridgeModelSvc.Delete(r.Context(), ctrModelId)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}
