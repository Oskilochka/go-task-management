package main

import (
	"josk/task-management-system/database"
	"josk/task-management-system/routes"
)

func main() {
	database.InitDB()

	routes.InitRoutes()
}
