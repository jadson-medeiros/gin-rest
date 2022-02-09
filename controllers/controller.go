package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jadson-medeiros/gin-rest/models"
)

func ShowAllStudents(c *gin.Context) {
	c.JSON(200, models.Students)
}

func Welcome(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"API:": "Hy " + name + ", welcome :D",
	})
}
