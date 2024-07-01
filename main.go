package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AddEmployee(w http.ResponseWriter, r *http.Request) {

}
func main() {
	db := initDB()
	h := newHandler(db)
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/Employees", h.postEmployee)
	r.Delete("/Employees", h.deleteEmployeesById)
	r.Post("/ListEmployeesByCompany", h.getEmployeesByCompanyId)
	r.Post("/ListEmployeesByComDep", h.getEmployeesByComDepId)
	r.Put("/UpdateUserByID", h.updateUserByID)
	log.Println("Microservice is running")
	log.Fatal(http.ListenAndServe("0.0.0.0:8888", r))

}
