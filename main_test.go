package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestSearchStudentByIDHandler(t *testing.T) {
	database.ConnectionDB()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.GET("/students/:id", controllers.SearchStudentByID)

	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	var studentMock models.Student

	json.Unmarshal(res.Body.Bytes(), &studentMock)

	assert.Equal(t, "Test", studentMock.Name, "The names should match")
	assert.Equal(t, "12345678901", studentMock.CPG)
	assert.Equal(t, http.StatusOK, res.Code)
}

func TestDeleteStudentHandler(t *testing.T) {
	database.ConnectionDB()

	CreateStudentMock()

	r := SetupTestRoutes()
	r.DELETE("/students/:id", controllers.DeleteStudent)

	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestUpdateStudentHandler(t *testing.T) {
	database.ConnectionDB()

	CreateStudentMock()
	defer DeleteStudentMock()

	r := SetupTestRoutes()
	r.PATCH("/students/:id", controllers.EditStudent)

	student := models.Student{Name: "Test2", CPG: "46123456789"}

	studentJson, _ := json.Marshal(student)
	path := "/students/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", path, bytes.NewReader(studentJson))

	res := httptest.NewRecorder()
	r.ServeHTTP(res, req)

	var studentUpdated models.Student

	json.Unmarshal(res.Body.Bytes(), &studentUpdated)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "Test2", studentUpdated.Name, "The names should match")
	assert.Equal(t, "46123456789", studentUpdated.CPG, "The CPG should match")
}
