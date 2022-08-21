package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rodrigoadao/simple-go-gin-api/database"
	"github.com/rodrigoadao/simple-go-gin-api/models"
)

func GetAllStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudentById(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	database.DB.First(&student, id)
	if (models.Student{} == student) {
		c.JSON(http.StatusNotFound, "Student not found!")
		return
	}
	c.JSON(http.StatusOK, student)
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusOK, student)
}

func DeleteStudentById(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	database.DB.Delete(&student, id)
	c.JSON(http.StatusOK, student)

}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)
	c.JSON(http.StatusOK, student)
}

func SearchStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Param("cpf")
	database.DB.Where(&models.Student{CPF: cpf}).First(&student)

	if (models.Student{} == student) {
		c.JSON(http.StatusNotFound, "Student not found!")
		return
	}
	c.JSON(http.StatusOK, student)
}
