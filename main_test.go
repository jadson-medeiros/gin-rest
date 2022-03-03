package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jadson-medeiros/gin-rest/controllers"
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

	if res.Code != http.StatusOK {
		t.Fatalf("Status error: value recieve was %d but the expected were %d",
			res.Code, http.StatusOK)
	}
}
