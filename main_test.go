package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rodrigoadao/simple-go-gin-api/controllers"
	"github.com/rodrigoadao/simple-go-gin-api/database"
	"github.com/rodrigoadao/simple-go-gin-api/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRouteTest() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "Teste", CPF: "12345678902", RG: "123456789"}
	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestGetAllStudents(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupRouteTest()
	r.GET("/students", controllers.GetAllStudents)

	req, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetStudentByCPF(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupRouteTest()
	r.GET("/student/cpf/:cpf", controllers.GetStudentByCPF)

	req, _ := http.NewRequest("GET", "/student/cpf/12345678902", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	var student models.Student
	json.NewDecoder(response.Body).Decode(&student)

	assert.Equal(t, ID, int(student.ID))
}

func TestGetStudentById(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupRouteTest()
	r.GET("/student/:id", controllers.GetStudentById)

	path := "/student/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	var student models.Student
	json.NewDecoder(response.Body).Decode(&student)

	assert.Equal(t, ID, int(student.ID))
}

func TestDeleteStudent(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	r := SetupRouteTest()
	r.DELETE("/student/:id", controllers.DeleteStudentById)

	path := "/student/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)
}

func TestUpdateStudent(t *testing.T) {
	database.ConnectDatabase()
	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupRouteTest()
	r.PATCH("/student/:id", controllers.UpdateStudent)

	student := models.Student{Name: "DeleteTest", CPF: "01987654321", RG: "987654321"}
	studentData, _ := json.Marshal(student)

	path := "/student/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewBuffer(studentData))
	response := httptest.NewRecorder()
	r.ServeHTTP(response, req)

	var updatedStudent models.Student
	json.Unmarshal(response.Body.Bytes(), &updatedStudent)

	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, student.Name, updatedStudent.Name)
}
