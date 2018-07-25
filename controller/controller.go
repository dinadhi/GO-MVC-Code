package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"mvc/controller/employee"

)

// New .
func New() http.Handler {
	r := mux.NewRouter()


	r.HandleFunc("/employee", employee.New).Methods("POST")
	r.HandleFunc("/employee/{id}", employee.Get).Methods("GET")

	return r
}
