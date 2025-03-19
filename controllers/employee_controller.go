package controllers

import (
    "poc_golang/config"
    "poc_golang/models"
    "net/http"
    // "encoding/json"

    "github.com/gin-gonic/gin"
)
// Create Multiple Employees
func CreateEmployee(c *gin.Context) {
    var employees []models.Employee // Ensure this is a slice

    // Bind JSON array to employees slice
    if err := c.ShouldBindJSON(&employees); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Bulk insert employees
    result := config.DB.Create(&employees)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Employees added successfully", "employees": employees})
}


 

func GetEmployees(c *gin.Context){
	var employees []models.Employee
	config.DB.Find(&employees)
	c.JSON(http.StatusOK, employees)
}


func GetEmployee(c *gin.Context){
	var employee models.Employee
	id := c.Param("id")
	if err := config.DB.First(&employee, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error" : "Employee not found"})
		return
	}
	c.JSON(http.StatusOK, employee)
}


func UpdateEmployee(c *gin.Context){
	var employee models.Employee
	id := c.Param("id")
	if err := config.DB.First(&employee, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error" : "Employee not found"})
		return
	}
	if err := c.ShouldBindJSON(&employee); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	config.DB.Save(&employee)
	c.JSON(http.StatusOK, employee)
}


func DeleteEmployee(c *gin.Context){
	var employee models.Employee
	id := c.Param("id")
	if err := config.DB.First(&employee, id).Error; err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error" : "Employee not found"})
		return
	}
	config.DB.Delete(&employee)
	c.JSON(http.StatusOK, gin.H{"message" : "Deleted successfully"})
}