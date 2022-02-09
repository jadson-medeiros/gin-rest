package main

import "github.com/gin-gonic/gin"

func showAllStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Git Hub",
	})
}

func main() {
	r := gin.Default()
	r.GET("/students", showAllStudents)
	r.Run()
}
