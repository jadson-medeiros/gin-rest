package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jadson-medeiros/gin-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/students", controllers.ShowAllStudents)
	r.Run()
}
