package routes

import(
	"poc_golang/controllers"
	"github.com/gin-gonic/gin"
)


func EmployeeRoutes(r *gin.Engine){
	r.POST("/create_employee", controllers.CreateEmployee)
	r.GET("/get_employees", controllers.GetEmployees)
	r.GET("/get_employee/:id", controllers.GetEmployee)
	r.PUT("/update_employee/:id", controllers.UpdateEmployee)
	r.DELETE("/delete_employee/:id", controllers.DeleteEmployee)
}