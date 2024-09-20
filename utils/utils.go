package utils

import (
	"encoding/json"
	"josk/task-management-system/models"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return
	}
}

func MapTaskToTaskResponse(task models.Task) models.TaskResponse {
	return models.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func MapTasksToTaskResponses(tasks []models.Task) []models.TaskResponse {
	taskResponses := make([]models.TaskResponse, len(tasks))

	for i, task := range tasks {
		taskResponses[i] = MapTaskToTaskResponse(task)
	}

	return taskResponses
}
