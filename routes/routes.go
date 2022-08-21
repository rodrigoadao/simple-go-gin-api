package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigoadao/simple-go-gin-api/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/student", controllers.GetAllStudents)
	r.GET("/student/:id", controllers.GetStudentById)
	r.GET("/student/cpf/:cpf", controllers.SearchStudentByCPF)
	r.POST("/student", controllers.CreateStudent)
	r.PATCH("/student/:id", controllers.UpdateStudent)
	r.DELETE("/student/:id", controllers.DeleteStudentById)
	r.Run()
}
