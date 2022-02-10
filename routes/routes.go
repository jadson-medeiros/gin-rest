package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jadson-medeiros/gin-rest/controllers"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/students", controllers.ShowAllStudents)
	r.GET("/:name", controllers.Welcome)
	r.GET("/students/:id", controllers.SearchStudentByID)

	r.POST("/students", controllers.CreateNewStudent)

	r.DELETE("/students/:id", controllers.DeleteStudent)

	r.PATCH("/students/:id", controllers.EditStudent)

	r.Run()
}
