package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"josk/task-management-system/database"
	"josk/task-management-system/middlewares"
	"josk/task-management-system/models"
	"josk/task-management-system/utils"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserKey).(*models.Claims)

	if !ok || claims == nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid token or context"}, http.StatusUnauthorized)
		return
	}

	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	task.UserID = claims.UserID

	result := database.DB.Create(&task)

	if result.Error != nil {
		utils.SendJSONResponse(w, map[string]string{"error": result.Error.Error()}, http.StatusInternalServerError)
		return
	}

	taskResponse := utils.MapTaskToTaskResponse(task)

	utils.SendJSONResponse(w, taskResponse, http.StatusCreated)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserKey).(*models.Claims)
	if !ok {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid token"}, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	var task models.Task

	result := database.DB.Where("id=? AND user_id = ?", id, claims.UserID).First(&task)

	if result.Error != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Task not found"}, http.StatusNotFound)
		return
	}

	var updatedTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)

	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid request payload"}, http.StatusBadRequest)
		return
	}

	database.DB.Model(&task).Updates(updatedTask)

	taskResponse := utils.MapTaskToTaskResponse(updatedTask)

	utils.SendJSONResponse(w, taskResponse, http.StatusOK)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserKey).(*models.Claims)

	if !ok || claims == nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid token or context"}, http.StatusUnauthorized)
		return
	}

	var tasks []models.Task

	result := database.DB.Where("user_id=?", claims.UserID).Find(&tasks)

	if result.Error != nil {
		utils.SendJSONResponse(w, map[string]string{"error": result.Error.Error()}, http.StatusInternalServerError)
		return
	}

	taskResponse := utils.MapTasksToTaskResponses(tasks)

	utils.SendJSONResponse(w, taskResponse, http.StatusOK)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserKey).(*models.Claims)
	if !ok {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid token"}, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	var task models.Task

	result := database.DB.Where("id=? AND user_id=?", id, claims.UserID).First(&task)

	if result.Error != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Task not found"}, http.StatusNotFound)
		return
	}

	taskResponse := utils.MapTaskToTaskResponse(task)

	utils.SendJSONResponse(w, taskResponse, http.StatusOK)
}

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	claims, ok := r.Context().Value(middlewares.UserKey).(*models.Claims)

	if !ok {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid token"}, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]

	result := database.DB.Where("id=? AND user_id = ?", id, claims.UserID).Delete(&models.Task{})

	if result.Error != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Task not found"}, http.StatusNotFound)
		return
	}

	utils.SendJSONResponse(w, map[string]string{"message": "Task was deleted"}, http.StatusOK)
}
