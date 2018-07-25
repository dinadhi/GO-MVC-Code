package employee

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"mvc/models/employee"
)

// New .
func New(w http.ResponseWriter, r *http.Request) {
	_employee, err := employee.New(r.FormValue("name"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "%s", _employee.ID.Hex())
	return
}

// Get .
func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_employee, err := employee.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if _employee == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, _employee.String())
	return
}
