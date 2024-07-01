package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type handler struct {
	db *gorm.DB
}

func newHandler(db *gorm.DB) handler {
	return handler{db}
}

// Create employee, return id
func (h handler) postEmployee(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var model Model
	err = json.Unmarshal(body, &model)
	if err != nil {
		log.Fatal(err)
	}

	// Create record in Employees
	employee := Employee{
		ID:           model.ID,
		Name:         model.Name,
		Surname:      model.Surname,
		Phone:        model.Phone,
		CompanyID:    model.CompanyID,
		DepartmentID: model.DepartmentID,
	}
	result := h.db.Create(&employee)
	if result.Error != nil {
		fmt.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("InternalServerError")
		return
	}

	// Create record in Passports
	passport := Passport{
		EmployeeID: model.Passport.EmployeeID,
		Employee:   model.Passport.Employee,
		Type:       model.Passport.Type,
		Number:     model.Passport.Number,
	}
	if result := h.db.Create(&passport); result.Error != nil {
		fmt.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("InternalServerError")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	// responceBody := map[string]string{"ID": fmt.Sprint(employee.ID), "Status": "Created"}
	responceBody := struct {
		ID     uint
		Status string
	}{
		ID:     employee.ID,
		Status: "Created",
	}
	json.NewEncoder(w).Encode(responceBody)
}

// Delete employee by ID
func (h handler) deleteEmployeesById(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var id struct {
		ID uint
	}
	err = json.Unmarshal(body, &id)
	if err != nil {
		log.Fatal(err)
	}

	if result := h.db.Where("employee_id = ?", id.ID).Delete(&Passport{}); result.Error != nil {
		fmt.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("InternalServerError")
		return
	}

	if result := h.db.Where("id = ?", id.ID).Delete(&Employee{}); result.Error != nil {
		fmt.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("InternalServerError")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responceBody := struct {
		ID     uint
		Status string
	}{
		ID:     id.ID,
		Status: "Deleted",
	}
	json.NewEncoder(w).Encode(responceBody)
}

// Get Employee fields by Company ID
func (h handler) getEmployeesByCompanyId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var companyid struct {
		CompanyID uint
	}
	err = json.Unmarshal(body, &companyid)
	if err != nil {
		log.Fatal(err)
	}
	var employees []Employee

	result := h.db.Where("company_id = ?", companyid.CompanyID).Find(&employees)
	if result.Error != nil {
		fmt.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("InternalServerError")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responceBody := employees
	json.NewEncoder(w).Encode(responceBody)
}

// Get Employee fields by Company ID
func (h handler) getEmployeesByComDepId(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var departmentid struct {
		DepartmentID uint
	}
	err = json.Unmarshal(body, &departmentid)
	if err != nil {
		log.Fatal(err)
	}
	var employees []Employee

	result := h.db.Where("department_id = ?", departmentid.DepartmentID).Find(&employees)
	if result.Error != nil {
		fmt.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("InternalServerError")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responceBody := employees
	json.NewEncoder(w).Encode(responceBody)
}

// Update Employee and Passport fields by ID
func (h handler) updateUserByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var updateData map[string]interface{}
	err = json.Unmarshal(body, &updateData)
	if err != nil {
		log.Fatal(err)
	}

	if _, ok := updateData["ID"]; !ok {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		err := map[string]string{
			"Error": "Can't get employee ID(or enter employee's ID into \"ID\" field)",
		}
		json.NewEncoder(w).Encode(err)
		return
	}

	if passport, ok := updateData["passport"]; ok {
		result := h.db.Model(&Passport{}).Where("employee_id = ?", updateData["ID"]).Updates(passport)
		if result.Error != nil {
			fmt.Println(result.Error)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode("InternalServerError")
			return
		}
		delete(updateData, "passport")
	}

	result := h.db.Model(&Employee{}).Where("id = ?", updateData["ID"]).Updates(updateData)
	if result.Error != nil {
		fmt.Println(result.Error)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("InternalServerError")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	responceBody := map[string]interface{}{"ID": updateData["ID"], "Status": "Updated"}
	json.NewEncoder(w).Encode(responceBody)
}
