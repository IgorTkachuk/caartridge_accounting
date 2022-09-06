package vndr

import (
	"encoding/json"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/vndr"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

const (
	vendorsURL = "/api/vendors"
	vendorURL  = "/api/vendors/:id"
)

var _ handlers.Handler = &Handler{}

type Handler struct {
	VendorService vndr.Service
}

func (h Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, vendorsURL, jwt.Middleware(apperror.Middleware(h.GetVendors)))
	router.HandlerFunc(http.MethodPost, vendorsURL, jwt.Middleware(apperror.Middleware(h.CreateVendor)))
	router.HandlerFunc(http.MethodGet, vendorURL, jwt.Middleware(apperror.Middleware(h.GetVendorById)))
	router.HandlerFunc(http.MethodDelete, vendorURL, jwt.Middleware(apperror.Middleware(h.DeleteVendor)))
	router.HandlerFunc(http.MethodPatch, vendorsURL, jwt.Middleware(apperror.Middleware(h.UpdateVendor)))
}

func (h *Handler) GetVendors(w http.ResponseWriter, r *http.Request) error {
	logger := logging.GetLogger()
	logger.Info("Execute Get all vendors handler")
	vendors, err := h.VendorService.GetAll(r.Context())
	if err != nil {
		return err
	}

	vendorsBytes, err := json.Marshal(vendors)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(vendorsBytes)

	return nil
}

func (h *Handler) CreateVendor(w http.ResponseWriter, r *http.Request) error {
	logger := logging.GetLogger()
	logger.Info("Creating new vendor")

	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()
	var dto vndr.CreateVendorDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		return apperror.BadRequestError("failed to decode data")
	}

	vendorId, err := h.VendorService.Create(r.Context(), dto)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("%s/%s", vendorsURL, vendorId))

	return nil
}

func (h *Handler) DeleteVendor(w http.ResponseWriter, r *http.Request) error {

	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	vendorID, _ := strconv.Atoi(params.ByName("id"))
	vendorDTO := vndr.DeleteVendorDTO{
		ID: vendorID,
	}

	_, err := h.VendorService.Delete(r.Context(), vendorDTO.ID)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h Handler) UpdateVendor(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var dto vndr.UpdateVendorDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		fmt.Println(err)
		return apperror.BadRequestError("failed to decode data")
	}

	err := h.VendorService.Update(r.Context(), dto)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}

func (h Handler) GetVendorById(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	vendorID, _ := strconv.Atoi(params.ByName("id"))

	v, err := h.VendorService.GetById(r.Context(), vendorID)
	if err != nil {
		return err
	}

	vendorBytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(vendorBytes)

	return nil
}
