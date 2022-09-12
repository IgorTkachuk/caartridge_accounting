package decom_cause

import (
	"encoding/json"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/decom_cause"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

const (
	decomcausesUrl = "/api/decomcause"
	decomcauseUrl  = "/api/decomcause/:id"
)

var _ handlers.Handler = &Handler{}

type Handler struct {
	DecomCauseSvc decom_cause.Service
}

func (h Handler) Register(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, decomcausesUrl, jwt.Middleware(apperror.Middleware(h.GetAllDecomCause)))
	r.HandlerFunc(http.MethodGet, decomcauseUrl, jwt.Middleware(apperror.Middleware(h.GetByIdDecomCause)))
	r.HandlerFunc(http.MethodPost, decomcausesUrl, jwt.Middleware(apperror.Middleware(h.CreateDecomCause)))
	r.HandlerFunc(http.MethodPatch, decomcausesUrl, jwt.Middleware(apperror.Middleware(h.UpdateDecomCause)))
	r.HandlerFunc(http.MethodDelete, decomcauseUrl, jwt.Middleware(apperror.Middleware(h.DeleteDecomCause)))
}

func (h Handler) GetAllDecomCause(w http.ResponseWriter, r *http.Request) error {
	decomCuses, err := h.DecomCauseSvc.GetAll(r.Context())
	if err != nil {
		return err
	}

	decomCausesBytes, err := json.Marshal(decomCuses)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(decomCausesBytes)

	return nil
}

func (h Handler) GetByIdDecomCause(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	decomCause, err := h.DecomCauseSvc.GetById(r.Context(), id)
	if err != nil {
		return err
	}

	decomCauseBytes, err := json.Marshal(decomCause)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(decomCauseBytes)

	return nil
}

func (h Handler) CreateDecomCause(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var dto decom_cause.CreateDecomCauseDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return err
	}

	id, err := h.DecomCauseSvc.Create(r.Context(), dto)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("%s/%s", decomcausesUrl, id))

	return nil
}

func (h Handler) UpdateDecomCause(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var dto decom_cause.UpdateDecomCauseDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return err
	}

	err = h.DecomCauseSvc.Update(r.Context(), dto)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h Handler) DeleteDecomCause(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	err = h.DecomCauseSvc.Delete(r.Context(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
