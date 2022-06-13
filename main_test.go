package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jadson-medeiros/gin-rest/controllers"
	"github.com/jadson-medeiros/gin-rest/database"
	"github.com/jadson-medeiros/gin-rest/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()

	return routes
}

func CreateStudentMock() {
	student := models.Student{Name: "Test", CPG: "12345678901"}

	database.DB.Create(&student)
	ID = int(student.ID)
}

func DeleteStudentMock() {
	var student models.Student
	database.DB.Delete(&student, ID)
}

func TestCheckStatusCodeWelcomeWithParams(t *testing.T) {
	r := SetupTestRoutes()

	r.GET("/:name", controllers.Welcome)

	req, _ := http.NewRequest("GET", "/test", nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code, "should be equals")

	mock := `{"API:":"Hi test, welcome :D"}`

	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, mock, string(resBody))
}

func TestGetAllStudentesHandler(t *testing.T) {
	database.ConnectionDB()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()

	r.GET("/students", controllers.ShowAllStudents)

	req, _ := http.NewRequest("GET", "/students", nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestSearchStudentByCPGHandler(t *testing.T) {
	database.ConnectionDB()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/students/cpg/:cpg", controllers.SearchStudentByCPG)

	req, _ := http.NewRequest("GET", "/students/cpg/12345678901", nil)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
