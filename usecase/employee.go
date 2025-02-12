package usecase

import (
	"encoding/json"
	"go-rest-api/model"
	"go-rest-api/repository"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *EmployeeService) CreateEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res:= &Response{}

	defer json.NewEncoder(w).Encode(res)
	var emp model.Employee

	err:= json.NewDecoder(r.Body).Decode(&emp)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error decoding request body: ", err)
		res.Error = err.Error()
		return
	}

	// assign new emp id
	emp.EmployeeID = uuid.NewString()
	repo:= repository.EmployeeRepo{MongoCollection: svc.MongoCollection}	

	// Insert new employee
	insertID, err:= repo.InsertEmployee(&emp)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error inserting employee: ", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp.EmployeeID
	w.WriteHeader(http.StatusOK)

	log.Println("Employee inserted with ID: ", insertID)
}
func (svc *EmployeeService) GetEmployeeByID(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res:= &Response{}
	defer json.NewEncoder(w).Encode(res)

	// Get Employee ID from request
	empID:= mux.Vars(r)["id"]
	log.Println("Employee ID: ", empID)

	repo:= repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err:= repo.FindEmployeeByID(empID)

	if err != nil{	
		w.WriteHeader(http.StatusNotFound)
		log.Println("Error getting employee: ", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)

	log.Println("Employee found: ", emp)
}
func (svc *EmployeeService) GetAllEmployees(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res:= &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo:= repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	emp, err:= repo.FindAllEmployee()

	if err != nil{	
		w.WriteHeader(http.StatusNotFound)
		log.Println("Error getting employee: ", err)
		res.Error = err.Error()
		return
	}

	res.Data = emp
	w.WriteHeader(http.StatusOK)

	log.Println("Employee found: ", emp)
}
func (svc *EmployeeService) UpdateEmployeeByID(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res:= &Response{}
	defer json.NewEncoder(w).Encode(res)

	// Get Employee ID from request
	empID:= mux.Vars(r)["id"]

	if empID == ""{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Employee ID is empty")
		res.Error = "Employee ID is empty"
		return
	}

	// Get updated employee from body
	var emp model.Employee

	err:= json.NewDecoder(r.Body).Decode(&emp)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error decoding request body: ", err)
		res.Error = err.Error()
		return
	}

	emp.EmployeeID = empID

	repo:= repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Update employee
	modifiedCount, err:= repo.UpdateEmployeeByID(empID, &emp)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error updating employee: ", err)
		res.Error = err.Error()
		return
	}

	res.Data = modifiedCount
	w.WriteHeader(http.StatusOK)

	log.Println("Employee updated with ID: ", empID)
}
func (svc *EmployeeService) DeleteEmployeeByID(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res:= &Response{}
	defer json.NewEncoder(w).Encode(res)

	// Get Employee ID from request
	empID:= mux.Vars(r)["id"]
	log.Println("Employee ID: ", empID)

	repo:= repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Delete employee
	deletedCount, err:= repo.DeleteEmployeeByID(empID)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error deleting employee: ", err)
		res.Error = err.Error()
		return
	}

	res.Data = deletedCount
	w.WriteHeader(http.StatusOK)

	log.Println("Employee deleted with ID: ", empID)
}
func (svc *EmployeeService) DeleteAllEmployee(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	res:= &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo:= repository.EmployeeRepo{MongoCollection: svc.MongoCollection}

	// Delete all employees
	deletedCount, err:= repo.DeleteAllEmployee()

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error deleting employee: ", err)
		res.Error = err.Error()
		return
	}

	res.Data = deletedCount
	w.WriteHeader(http.StatusOK)

	log.Println("All employees deleted")
}