package ctr_showcase

import (
	"encoding/json"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/ctr_showcase"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	ctrShowcaseUrl = "/api/ctrshowcase"
)

var _ handlers.Handler = &Handler{}

type Handler struct {
	CtrShowcaseSvc ctr_showcase.Service
}

func (h Handler) Register(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, ctrShowcaseUrl, jwt.Middleware(apperror.Middleware(h.GetCtrShowcase)))
}

func (h Handler) GetCtrShowcase(w http.ResponseWriter, r *http.Request) error {
	ctrs, err := h.CtrShowcaseSvc.GetAll(r.Context())
	if err != nil {
		return err
	}

	ctrsBytes, err := json.Marshal(ctrs)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(ctrsBytes)

	return nil
}
