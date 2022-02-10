package main

import (
	"github.com/jadson-medeiros/gin-rest/database"
	"github.com/jadson-medeiros/gin-rest/routes"
)

func main() {
	database.ConnectionDB()

	routes.HandleRequests()
}
