package controllers

import "github.com/gin-gonic/gin"

func ShowAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Git Hub",
	})
}

func Welcome(c *gin.Context) {
	name := c.Params.ByName("name")

	c.JSON(200, gin.H{
		"API:": "Hy " + name + ", welcome :D",
	})
}
