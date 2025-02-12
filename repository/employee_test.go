package repository

import (
	"context"
	"go-rest-api/model"
	"log"
	"testing"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client{
	mongoTestClient, err:= mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://abhisawarkar85:ZOYbmXQ9nEsi6fjU@cluster0.y3vlf.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	if err != nil{
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	log.Println("Connected to MongoDB")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())

	if err != nil{
		log.Fatal("Error pinging to MongoDB: ", err)
	}

	log.Println("Ping to MongoDB successful")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T){
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	// Dummy date
	emp1:= uuid.New().String()
	emp2:= uuid.New().String()

	// Connect to collection
	coll:= mongoTestClient.Database("companyDb").Collection("employees_test")

	empRepo:= EmployeeRepo{MongoCollection: coll}

	// Insert employee 1 data
	t.Run("Insert Employee 1", func(t *testing.T){
		emp:= model.Employee{
			Name: "Abhinav Sawarkar",
			Department: "IBM",
			EmployeeID: emp1,
		}

		result, err:= empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("Insert 1 operation failed", err)
		}

		t.Log("Insert 1 operation successful", result)
	})
	// Insert employee 2 data
	t.Run("Insert Employee 1", func(t *testing.T){
		emp:= model.Employee{
			Name: "Manas Sawarkar",
			Department: "InternPro",
			EmployeeID: emp2,
		}

		result, err:= empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("Insert 1 operation failed", err)
		}

		t.Log("Insert 1 operation successful", result)
	})

	// Get Employee by ID
	t.Run("Get Employee 1", func(t *testing.T){
		result, err:= empRepo.FindEmployeeByID(emp1)

		if err != nil{
			t.Fatal("Get operations failed", err)
		}

		t.Log("emp1", result.Name)
	})

	// Get All Employee Data
	t.Run("Get all employee", func(t *testing.T){
		result, err:= empRepo.FindAllEmployee()

		if err != nil{
			t.Fatal("Get operations failed", err)
		}

		t.Log("employees", result)
	})

	// Update Employee Data
	t.Run("Update Employee 1", func(t *testing.T){
		emp:= model.Employee{
			Name: "Abhinav Kiran Sawarkar",
			Department: "Physics",
			EmployeeID: emp1,
		}

		result, err:= empRepo.UpdateEmployeeByID(emp1, &emp)

		if err != nil{
			t.Fatal("Update operation failed", err)
		}

		t.Log("Update operation successful", result)
	})

	// Get Employee by ID after update
	t.Run("Get Employee 1", func(t *testing.T){
		result, err:= empRepo.FindEmployeeByID(emp1)

		if err != nil{
			t.Fatal("Get operations failed", err)
		}

		t.Log("emp1", result.Name)
	})

	// Delete Employee by ID
	t.Run("Delete Employee 1", func(t *testing.T){
		result, err:= empRepo.DeleteEmployeeByID(emp1)

		if err != nil{
			t.Fatal("Delete operation failed", err)
		}

		t.Log("Delete operation successful", result)
	})

	// Get All Employee Data
	t.Run("Get all employee", func(t *testing.T){
		result, err:= empRepo.FindAllEmployee()

		if err != nil{
			t.Fatal("Get operations failed", err)
		}

		t.Log("employees", result)
	})

	// Delete All Employee Data
	t.Run("Delete all employee", func(t *testing.T){
		result, err:= empRepo.DeleteAllEmployee()

		if err != nil{
			t.Fatal("Delete operations failed", err)
		}

		t.Log("Delete all operation successful", result)
	})
}