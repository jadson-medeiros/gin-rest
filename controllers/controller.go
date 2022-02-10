package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jadson-medeiros/gin-rest/database"
	"github.com/jadson-medeiros/gin-rest/models"
)

func ShowAllStudents(c *gin.Context) {
	var students []models.Student

	database.DB.Find(&students)

	c.JSON(200, students)
}

func Welcome(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"API:": "Hy " + name + ", welcome :D",
	})
}

func CreateNewStudent(c *gin.Context) {
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

func SearchStudentByID(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}
