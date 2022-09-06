package bisiness_line

import (
	"encoding/json"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/business_line"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

var _ handlers.Handler = &Handler{}

const (
	blsUrl = "/api/bl"
	blUrl  = "/api/bl/:id"
)

type Handler struct {
	BusinessLineSvc business_line.Service
}

func (h Handler) Register(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, blsUrl, jwt.Middleware(apperror.Middleware(h.GetBL)))
	r.HandlerFunc(http.MethodGet, blUrl, jwt.Middleware(apperror.Middleware(h.GetBLById)))
	r.HandlerFunc(http.MethodPost, blsUrl, jwt.Middleware(apperror.Middleware(h.CreateBL)))
	r.HandlerFunc(http.MethodPatch, blsUrl, jwt.Middleware(apperror.Middleware(h.UpdateBL)))
	r.HandlerFunc(http.MethodDelete, blUrl, jwt.Middleware(apperror.Middleware(h.DeleteBL)))
}

func (h Handler) GetBL(w http.ResponseWriter, r *http.Request) error {
	bls, err := h.BusinessLineSvc.GetAll(r.Context())
	if err != nil {
		return err
	}

	blsBytes, err := json.Marshal(bls)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(blsBytes)

	return nil
}

func (h Handler) GetBLById(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	bl, err := h.BusinessLineSvc.GetById(r.Context(), id)
	if err != nil {
		return err
	}

	blBytes, err := json.Marshal(bl)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(blBytes)

	return nil
}

func (h Handler) CreateBL(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var bl business_line.CreateBusinessLineDTO

	if err := json.NewDecoder(r.Body).Decode(&bl); err != nil {
		appErr := apperror.BadRequestError("Failed to decode info for new Business Line entity")
		return appErr
	}

	id, err := h.BusinessLineSvc.Create(r.Context(), bl)
	if err != nil {
		return err
	}

	w.Header().Set("Location", fmt.Sprintf("%s/%s", blsUrl, id))
	w.WriteHeader(http.StatusCreated)

	return nil
}

func (h Handler) UpdateBL(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var bl business_line.UpdateBusinessLineDTO

	if err := json.NewDecoder(r.Body).Decode(&bl); err != nil {
		appErr := apperror.BadRequestError("Failed to decode info for update Business Line entity")
		return appErr
	}

	err := h.BusinessLineSvc.Update(r.Context(), bl)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}

func (h Handler) DeleteBL(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	err = h.BusinessLineSvc.Delete(r.Context(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}


