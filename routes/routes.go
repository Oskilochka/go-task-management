package routes

import (
	"github.com/gorilla/mux"
	"josk/task-management-system/handlers"
	"josk/task-management-system/middlewares"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	// auth
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// tasks
	taskRoutes := router.PathPrefix("/tasks").Subrouter()
	taskRoutes.Use(middlewares.JWTMiddleware)
	taskRoutes.HandleFunc("", handlers.CreateTask).Methods("POST")
	taskRoutes.HandleFunc("", handlers.GetTasks).Methods("GET")
	taskRoutes.HandleFunc("/{id}", handlers.GetTaskById).Methods("GET")
	taskRoutes.HandleFunc("/{id}", handlers.UpdateTask).Methods("PUT")
	taskRoutes.HandleFunc("/{id}", handlers.DeleteTaskById).Methods("DELETE")

	return router
}
