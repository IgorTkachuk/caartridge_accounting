package ou

import (
	"encoding/json"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/ou"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

const (
	ousURL = "/api/ous"
	ouURL  = "/api/ous/:id"
)

var _ handlers.Handler = &Handler{}

type Handler struct {
	OuService ou.Service
}

func (h Handler) Register(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, ousURL, jwt.Middleware(apperror.Middleware(h.GetOus)))
	r.HandlerFunc(http.MethodPost, ousURL, jwt.Middleware(apperror.Middleware(h.CreateOu)))
	r.HandlerFunc(http.MethodGet, ouURL, jwt.Middleware(apperror.Middleware(h.GetOuById)))
	r.HandlerFunc(http.MethodDelete, ouURL, jwt.Middleware(apperror.Middleware(h.DeleteOu)))
	r.HandlerFunc(http.MethodPatch, ousURL, jwt.Middleware(apperror.Middleware(h.UpdateOu)))
}

func (h Handler) GetOus(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")

	ous, err := h.OuService.GetAll(r.Context())
	if err != nil {
		return err
	}

	ousBytes, err := json.Marshal(ous)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(ousBytes)

	return nil
}

func (h Handler) CreateOu(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var newOu ou.CreateOuDTO
	if err := json.NewDecoder(r.Body).Decode(&newOu); err != nil {
		return apperror.BadRequestError("Failed to decode new Ou data")
	}

	newOuId, err := h.OuService.Create(r.Context(), newOu)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Location", fmt.Sprintf("%s/%s", ousURL, newOuId))

	return nil
}

func (h Handler) GetOuById(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	ouId, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	ou, err := h.OuService.GetById(r.Context(), ouId)
	if err != nil {
		return err
	}

	ouBytes, err := json.Marshal(ou)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(ouBytes)

	return nil
}

func (h Handler) DeleteOu(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	ouId, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	err = h.OuService.Delete(r.Context(), ouId)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil

}

func (h Handler) UpdateOu(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var ou ou.UpdateOuDTO
	if err := json.NewDecoder(r.Body).Decode(&ou); err != nil {
		return apperror.BadRequestError("Failed to decode data for update OU")
	}

	err := h.OuService.Update(r.Context(), ou)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
