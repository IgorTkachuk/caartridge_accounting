package employee

import (
	"encoding/json"
	"fmt"
	"github.com/IgorTkachuk/cartridge_accounting/internal/apperror"
	"github.com/IgorTkachuk/cartridge_accounting/internal/domain/employee"
	"github.com/IgorTkachuk/cartridge_accounting/internal/handlers"
	"github.com/IgorTkachuk/cartridge_accounting/pkg/jwt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

const (
	employeesUrl = "/api/employee"
	employeeUrl  = "/api/employee/:id"
)

var _ handlers.Handler = &Handler{}

type Handler struct {
	EmployeeService employee.Service
}

func (h Handler) Register(r *httprouter.Router) {
	r.HandlerFunc(http.MethodGet, employeesUrl, jwt.Middleware(apperror.Middleware(h.GetEmployees)))
	r.HandlerFunc(http.MethodGet, employeeUrl, jwt.Middleware(apperror.Middleware(h.GetByIdEmployees)))
	r.HandlerFunc(http.MethodPost, employeesUrl, jwt.Middleware(apperror.Middleware(h.CreateEmployee)))
	r.HandlerFunc(http.MethodPatch, employeesUrl, jwt.Middleware(apperror.Middleware(h.UpdateEmployee)))
	r.HandlerFunc(http.MethodDelete, employeeUrl, jwt.Middleware(apperror.Middleware(h.DeleteEmployee)))
}

func (h Handler) GetEmployees(w http.ResponseWriter, r *http.Request) error {
	es, err := h.EmployeeService.GetAll(r.Context())
	if err != nil {
		return err
	}

	esBytes, err := json.Marshal(es)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(esBytes)

	return nil
}

func (h Handler) GetByIdEmployees(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	e, err := h.EmployeeService.GetById(r.Context(), id)
	if err != nil {
		return err
	}

	eBytes, err := json.Marshal(e)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(eBytes)

	return nil
}

func (h Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var e employee.CreateEmployeeDTO
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		newErr := apperror.BadRequestError("Failed to decode create new employee entity info")
		return newErr
	}

	id, err := h.EmployeeService.Create(r.Context(), e)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", fmt.Sprintf("%s/%s", employeesUrl, id))

	return nil
}

func (h Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	var e employee.UpdateEmployeeDTO
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		return err
	}

	err = h.EmployeeService.Update(r.Context(), e)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}

func (h Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) error {
	params := r.Context().Value(httprouter.ParamsKey).(httprouter.Params)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		return err
	}

	err = h.EmployeeService.Delete(r.Context(), id)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)

	return nil
}
