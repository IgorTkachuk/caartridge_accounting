package ctr_status_type

import (
	"encoding/json"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/cartridge_status_type"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

const (
	ctrstatustypesUrl = "/api/ctrstatustype"
	ctrstatustypeUrl  = "/api/ctrstatustype/:id"
)

var _ handlers.Handler = &Handler{}

type Handler struct {
	CartridgeStatusTypeSvc cartridge_status_type.Service
}

func (h Handler) Register(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, ctrstatustypesUrl, jwt.Middleware(apperror.Middleware(h.GetAllCtrStatusType)))
	r.HandlerFunc(http.MethodGet, ctrstatustypeUrl, jwt.Middleware(apperror.Middleware(h.GetByIDCtrStatusType)))
	r.HandlerFunc(http.MethodPost, ctrstatustypesUrl, jwt.Middleware(apperror.Middleware(h.CreateCtrStatusType)))
	r.HandlerFunc(http.MethodPatch, ctrstatustypesUrl, jwt.Middleware(apperror.Middleware(h.UpdateCtrStatusType)))
	r.HandlerFunc(http.MethodDelete, ctrstatustypeUrl, jwt.Middleware(apperror.Middleware(h.DeleteCtrStatusType)))
}

func (h Handler) GetAllCtrStatusType(w http.ResponseWriter, r *http.Request) error {
	cStatuses, err := h.CartridgeStatusTypeSvc.GetAll(r.Context())
	if err != nil {
		return err
	}

	cStatusesBytes, err := json.Marshal(cStatuses)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(cStatusesBytes)

	return nil
}

func (h Handler) GetByIDCtrStatusType(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	cStatusType, err := h.CartridgeStatusTypeSvc.GetById(r.Context(), id)
	if err != nil {
		return err
	}

	cStatusTypeBytes, err := json.Marshal(cStatusType)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(cStatusTypeBytes)

	return nil
}

func (h Handler) CreateCtrStatusType(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var dto cartridge_status_type.CreateCartridgeStatusTypeDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return apperror.BadRequestError("Failed decode data for new cartridge status type")
	}

	id, err := h.CartridgeStatusTypeSvc.Create(r.Context(), dto)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("%s/%s", ctrstatustypesUrl, id))

	return nil
}

func (h Handler) UpdateCtrStatusType(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var dto cartridge_status_type.UpdateCartridgeStatusTypeDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		return apperror.BadRequestError("Failed to decode data for update cartridge status type")
	}

	err = h.CartridgeStatusTypeSvc.Update(r.Context(), dto)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h Handler) DeleteCtrStatusType(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	err = h.CartridgeStatusTypeSvc.Delete(r.Context(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}
