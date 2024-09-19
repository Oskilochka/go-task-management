package main

import (
	"josk/task-management-system/database"
	"josk/task-management-system/routes"
	"log"
	"net/http"
)

func main() {
	database.InitDB()

	router := routes.SetupRouter()

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
}
