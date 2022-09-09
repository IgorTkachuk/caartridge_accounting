package doc_type

import (
	"encoding/json"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/doc_type"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

const (
	doctypesUrl = "/api/doctype"
	doctypeUrl  = "/api/doctype/:id"
)

var _ handlers.Handler = &Handler{}

type Handler struct {
	DocTypeService doc_type.Service
}

func (h Handler) Register(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, doctypesUrl, jwt.Middleware(apperror.Middleware(h.GetDocTypes)))
	r.HandlerFunc(http.MethodGet, doctypeUrl, jwt.Middleware(apperror.Middleware(h.GetByIdDocTypes)))
	r.HandlerFunc(http.MethodPost, doctypesUrl, jwt.Middleware(apperror.Middleware(h.CreateDocType)))
	r.HandlerFunc(http.MethodPatch, doctypesUrl, jwt.Middleware(apperror.Middleware(h.UpdateDocType)))
	r.HandlerFunc(http.MethodDelete, doctypeUrl, jwt.Middleware(apperror.Middleware(h.DeleteDocType)))
}

func (h Handler) GetDocTypes(w http.ResponseWriter, r *http.Request) error {
	docTypes, err := h.DocTypeService.FindAll(r.Context())
	if err != nil {
		return err
	}

	docTypesBytes, err := json.Marshal(docTypes)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(docTypesBytes)

	return nil
}

func (h Handler) GetByIdDocTypes(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	docType, err := h.DocTypeService.FindById(r.Context(), id)
	if err != nil {
		return err
	}

	docTypeByte, err := json.Marshal(docType)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(docTypeByte)

	return nil
}

func (h Handler) CreateDocType(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var docType doc_type.CreateDocTypeDTO
	if err := json.NewDecoder(r.Body).Decode(&docType); err != nil {
		newErr := apperror.BadRequestError("Fail to decode new doc type info due create entity")
		return newErr
	}

	id, err := h.DocTypeService.Create(r.Context(), docType)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("%s/%s", doctypesUrl, id))

	return nil
}

func (h Handler) UpdateDocType(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var docType doc_type.UpdateDocTypeDTO
	if err := json.NewDecoder(r.Body).Decode(&docType); err != nil {
		newErr := apperror.BadRequestError("Failed to decode doc type info due update entity")
		return newErr
	}

	err := h.DocTypeService.Update(r.Context(), docType)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (h Handler) DeleteDocType(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	err = h.DocTypeService.Delete(r.Context(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
