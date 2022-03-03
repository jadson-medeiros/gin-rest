package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jadson-medeiros/gin-rest/controllers"
	"github.com/jadson-medeiros/gin-rest/database"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()

	return routes
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

	r := SetupTestRoutes()

	r.GET("/students", controllers.ShowAllStudents)

	req, _ := http.NewRequest("GET", "/students", nil)

	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}
