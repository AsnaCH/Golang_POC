package controllers

import (
    "poc_golang/config"
    "poc_golang/models"
    "net/http"
    // "encoding/json" 
	// "fmt"
    "github.com/gin-gonic/gin"
)
func CreateEmployee(c *gin.Context) {
    var employees []models.Employee // Ensure this is a slice

    // Bind JSON array to employees slice
    if err := c.ShouldBindJSON(&employees); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Variables to store the results
    var existing_employees []string
    var newEmployees []models.Employee

    // Check for existing emails in the DB
    for _, emp := range employees {
        var existingEmployee models.Employee
        // Check if the employee exists by email or phone number
        if err := config.DB.Where("email = ? OR phoneno = ?", emp.Email, emp.Phoneno).First(&existingEmployee).Error; err == nil {
			entry := emp.Email + " with this phonenumber " + emp.Phoneno
			existing_employees = append(existing_employees, entry)
        } else {
            // If employee doesn't exist, add to newEmployees slice
            newEmployees = append(newEmployees, emp)
        }
    }

    // Prepare response message
    var responseMessages []gin.H

    // If any employees already exist, add their emails to the response
    if len(existing_employees) > 0 {
        for _, email := range existing_employees {
            responseMessages = append(responseMessages, gin.H{
                
                "message": "Employee with " + email + " already exists",
            })
        }
    }

    // Bulk insert new employees
    if len(newEmployees) > 0 {
        result := config.DB.Create(&newEmployees)
        if result.Error != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
            return
        }
        for _, emp := range newEmployees {
            responseMessages = append(responseMessages, gin.H{
                "email":   emp.Email,
                "message": "Employee with " + emp.Email + " has been created successfully",
            })
        }
    }

    // Return final response with messages about both existing and new employees
    c.JSON(http.StatusOK, gin.H{
        "messages": responseMessages,
    })
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