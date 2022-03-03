package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jadson-medeiros/gin-rest/controllers"
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
}
