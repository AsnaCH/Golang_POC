package main


import(
	"poc_golang/config"
	"poc_golang/models"
	"poc_golang/routes"
	"github.com/gin-gonic/gin"
)



func main(){
	config.ConnectDB()                                // Connect to database

	config.DB.AutoMigrate(&models.Employee{})         // Database Migration

	router := gin.Default()                           // Inintialize Gin router

	routes.EmployeeRoutes(router)                     // Setup routes

	router.Run(":8080")                               // Start the server


}