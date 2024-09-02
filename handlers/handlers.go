package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"josk/task-management-system/database"
	"josk/task-management-system/models"
	"josk/task-management-system/utils"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	result := database.DB.Create(&task)

	if result.Error != nil {
		utils.SendJSONResponse(w, map[string]string{"error": result.Error.Error()}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, task, http.StatusCreated)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var task models.Task
	result := database.DB.Find(&task, id)

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
	utils.SendJSONResponse(w, task, http.StatusOK)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task

	result := database.DB.Find(&tasks)

	if result.Error != nil {
		utils.SendJSONResponse(w, map[string]string{"error": result.Error.Error()}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, tasks, http.StatusOK)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var task models.Task

	result := database.DB.First(&task, id)
	if result.Error != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Task not found"}, http.StatusNotFound)
		return
	}

	utils.SendJSONResponse(w, task, http.StatusOK)
}

func DeleteTaskById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var task models.Task

	result := database.DB.First(&task, id)
	if result.Error != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Task not found"}, http.StatusNotFound)
		return
	}
	database.DB.Delete(&task)

	utils.SendJSONResponse(w, map[string]string{"message": "Task was deleted"}, http.StatusOK)
}
