package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"poc_golang/config"
	"poc_golang/models"
	"testing"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB(){
	var err error
	config.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil{
		panic("Failed to connect database")
	}
	config.DB.AutoMigrate(&models.Employee{})
}


func SetupRouter() *gin.Engine{
	r := gin.Default()
	r.POST("/create_employee", CreateEmployee)
	r.GET("/get_employees", GetEmployees)
	r.GET("/get_employee", GetEmployee)
	r.PUT("/update_employee", UpdateEmployee)
	r.DELETE("/delete_employee", DeleteEmployee)
	return r
}


// Run setup before tests
func TestMain(m *testing.M){
	SetupTestDB()
	code := m.Run()
	os.Exit(code)
}


func TestCreateEmployee(t *testing.T){
	r := SetupRouter()
	employee_data := `[{"name" : "Tom", "email" : "tom123@gmail.com", "position" : "Manager", "age" : 45, "salary" : 80000}]`
	req,_ := http.NewRequest("POST", "/create_employee", bytes.NewBuffer([]byte(employee_data)))
	req.Header.Set("Content-Type","application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK{
		t.Errorf("Expected status 200 got %d", w.Code)
	}
}

func TestGetEmployees(t *testing.T){
	r := SetupRouter()
	req,_ := http.NewRequest("GET", "/get_employees", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK{
		t.Errorf("Expected status 200 got %d", w.Code)
	}
}


func TestGetEmployee(t *testing.T){
	r := SetupRouter()
	employee := models.Employee{Name : "Alice", Email : "alice@example.com", Position : "HR", Age : 30, Salary : 50000}
	config.DB.Create(&employee)
	req,_ := http.NewRequest("GET", "/get_employee", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK{
		t.Errorf("Expected status 200 got %d", w.Code)
	}
}

func TestUpdateEmployee(t *testing.T){
	r := SetupRouter()
	employee := models.Employee{Name : "Bob", Email : "bob@example.com", Position : "Principal Engineer", Age : 50, Salary : 80000}
	config.DB.Create(&employee)
	updated_data := `{"name" : "Bobupdated", "email" : "bob@example.com", "position" : "Principal Engineer", "age" : 50, "salary" : 80000}`
	req,_ := http.NewRequest("PUT", "/update_employee", bytes.NewBuffer([]byte(updated_data)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK{
		t.Errorf("Expected status 200 got %d", w.Code)
	}
}


func TestDeleteEmployee(t *testing.T){
	r := SetupRouter()
	employee := models.Employee{Name : "John", Email : "john@example.com", Position : "Manager", Age : 49, Salary : 90000}
	config.DB.Create(&employee)
	req,_ := http.NewRequest("DELETE", "/delete_employee", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK{
		t.Errorf("Expected status 200 got %d", w.Code)
	}
}

