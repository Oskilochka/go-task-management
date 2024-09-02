package routes

import (
	"github.com/gorilla/mux"
	"josk/task-management-system/handlers"
	"log"
	"net/http"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetTaskById).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTaskById).Methods("DELETE")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		log.Fatalln("There's an error with the server,", err)
	}
	return router
}
