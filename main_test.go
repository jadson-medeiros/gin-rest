package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupTestRoutes() *gin.Engine {
	routes := gin.Default()

	return routes
}

func TestFail(t *testing.T) {
	t.Fatalf("Test failured, but it's ok!")
}
