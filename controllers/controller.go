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
		"API:": "Hi " + name + ", welcome :D",
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

	if err := models.ValidateData(&student); err != nil {
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

func DeleteStudent(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")

	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Student deleted with success",
	})
}

func EditStudent(c *gin.Context) {
	var student models.Student

	id := c.Params.ByName("id")

	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := models.ValidateData(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	c.JSON(http.StatusOK, student)
}

func SearchStudentByCPG(c *gin.Context) {
	var student models.Student

	cpg := c.Param("cpg")

	database.DB.Where(&models.Student{CPG: cpg}).First(&student)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Student not found",
		})

		return
	}

	c.JSON(http.StatusOK, student)
}

func ShowIndexPage(c *gin.Context) {
	var students models.Student

	database.DB.Find(&students)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"students": students,
	})
}
