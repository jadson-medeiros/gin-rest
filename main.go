package main

import (
	"github.com/jadson-medeiros/gin-rest/database"
	"github.com/jadson-medeiros/gin-rest/models"
	"github.com/jadson-medeiros/gin-rest/routes"
)

func main() {
	database.ConnectionDB()

	models.Students = []models.Student{
		{Name: "Git Hub", CPG: "00000000000", ID: "234567522"},
		{Name: "Git Lab", CPG: "00000000001", ID: "234567523"},
	}

	routes.HandleRequests()
}
