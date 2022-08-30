package prnt

import (
	"encoding/json"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/prnt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

const (
	printersURL = "/api/printers"
	printerURL  = "/api/printers/:id"
)

var _ handlers.Handler = &Handler{}

type Handler struct {
	PrinterService prnt.Service
}

func (h Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, printersURL, jwt.Middleware(apperror.Middleware(h.GetPrinters)))
	router.HandlerFunc(http.MethodPost, printersURL, jwt.Middleware(apperror.Middleware(h.CreatePrinter)))
	router.HandlerFunc(http.MethodGet, printerURL, jwt.Middleware(apperror.Middleware(h.GetPrinterById)))
	router.HandlerFunc(http.MethodDelete, printerURL, jwt.Middleware(apperror.Middleware(h.DeletePrinter)))
	router.HandlerFunc(http.MethodPatch, printersURL, jwt.Middleware(apperror.Middleware(h.UpdatePrinter)))
}

func (h Handler) GetPrinters(w http.ResponseWriter, r *http.Request) error {
	logger := logging.GetLogger()
	logger.Info("Execute Get all printers handler")

	w.Header().Set("Content-Type", "application/json")

	printers, err := h.PrinterService.GetAll(r.Context())
	if err != nil {
		return err
	}

	printersBytes, err := json.Marshal(printers)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write(printersBytes)

	return nil
}

func (h Handler) CreatePrinter(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()
	var newPrinter prnt.CreatePrnDTO
	err := json.NewDecoder(r.Body).Decode(&newPrinter)
	if err != nil {
		return apperror.BadRequestError("Failed to decode new printer data")
	}
	printerId, err := h.PrinterService.Create(r.Context(), newPrinter)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("%s/%s", printersURL, printerId))
	return nil
}

func (h Handler) DeletePrinter(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	printerId, _ := strconv.Atoi(params.ByName("id"))

	printerDTO := prnt.DeletePrnDTO{
		ID: printerId,
	}

	_, err := h.PrinterService.Delete(r.Context(), printerDTO.ID)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h Handler) UpdatePrinter(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var updatePrnInfo prnt.UpdatePrnDTO

	if err := json.NewDecoder(r.Body).Decode(&updatePrnInfo); err != nil {
		return apperror.BadRequestError("failed to decode update printer info")
	}

	err := h.PrinterService.Update(r.Context(), updatePrnInfo)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h Handler) GetPrinterById(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	printerId, _ := strconv.Atoi(params.ByName("id"))

	printerInfo, err := h.PrinterService.GetById(r.Context(), printerId)
	if err != nil {
		return err
	}

	printerInfoBytes, err := json.Marshal(printerInfo)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(printerInfoBytes)

	return nil
}
