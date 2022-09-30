package doc

import (
	"encoding/json"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/doc"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

const (
	docsUrl = "/api/doc"
	docUrl  = "/api/doc/:id"
)

var _ handlers.Handler = &Handler{}

type Handler struct {
	DocSvc doc.Service
}

func (h Handler) Register(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, docsUrl, jwt.Middleware(apperror.Middleware(h.GetAllDoc)))
	r.HandlerFunc(http.MethodGet, docUrl, jwt.Middleware(apperror.Middleware(h.GetByIdDoc)))
	r.HandlerFunc(http.MethodPost, docsUrl, jwt.Middleware(apperror.Middleware(h.CreateDoc)))
	r.HandlerFunc(http.MethodPatch, docsUrl, jwt.Middleware(apperror.Middleware(h.UpdateDoc)))
	r.HandlerFunc(http.MethodDelete, docUrl, jwt.Middleware(apperror.Middleware(h.DeleteDoc)))
}

func (h Handler) GetAllDoc(w http.ResponseWriter, r *http.Request) error {
	docs, err := h.DocSvc.GetAll(r.Context())
	if err != nil {
		return err
	}

	docsBytes, err := json.Marshal(docs)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(docsBytes)

	return nil
}

func (h Handler) GetByIdDoc(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	doc, err := h.DocSvc.GetById(r.Context(), id)
	if err != nil {
		return err
	}

	docBytes, err := json.Marshal(doc)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(docBytes)

	return nil
}

func (h Handler) CreateDoc(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var dto doc.CreateDocDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return apperror.BadRequestError("Fail to decode doc data due parse incoming json")
	}

	userId, err := strconv.Atoi(r.Context().Value("user_id").(string))
	if err != nil {
		return err
	}

	dto.DocOwnerId = userId
	id, err := h.DocSvc.Create(r.Context(), dto)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("%s/%s", docsUrl, id))

	return nil
}

func (h Handler) UpdateDoc(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var dto doc.UpdateDocDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		return apperror.BadRequestError("Fail to decode doc data due parse incoming json for update")
	}

	err = h.DocSvc.Update(r.Context(), dto)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}

func (h Handler) DeleteDoc(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	err = h.DocSvc.Delete(r.Context(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}
