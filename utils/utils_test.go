package utils

import (
	"encoding/json"
	"gorm.io/gorm"
	"josk/task-management-system/models"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestSendJSONResponse(t *testing.T) {
	rr := httptest.NewRecorder()

	data := map[string]string{"message": "success"}

	SendJSONResponse(rr, data, http.StatusOK)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, status)
	}

	expectedContentType := "application/json"
	if rr.Header().Get("Content-Type") != expectedContentType {
		t.Errorf("Expected content type %s, but got %s", expectedContentType, rr.Header().Get("Content-Type"))
	}

	var responseData map[string]string
	err := json.NewDecoder(rr.Body).Decode(&responseData)
	if err != nil {
		t.Fatalf("Expected valid JSON response, but got error: %v", err)
	}

	if responseData["message"] != "success" {
		t.Errorf("Expected message to be 'success', but got %s", responseData["message"])
	}
}

func TestMapTaskToTaskResponse(t *testing.T) {
	task := models.Task{
		ID:          1,
		UserID:      1,
		Title:       "Test Task",
		Description: "This is a test task",
		Completed:   false,
		Model: gorm.Model{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	taskResponse := MapTaskToTaskResponse(task)

	if taskResponse.ID != task.ID {
		t.Errorf("Expected ID to be %d, but got %d", task.ID, taskResponse.ID)
	}
	if taskResponse.Title != task.Title {
		t.Errorf("Expected Title to be %s, but got %s", task.Title, taskResponse.Title)
	}
	if taskResponse.Description != task.Description {
		t.Errorf("Expected Description to be %s, but got %s", task.Description, taskResponse.Description)
	}
	if taskResponse.Completed != task.Completed {
		t.Errorf("Expected Completed to be %v, but got %v", task.Completed, taskResponse.Completed)
	}
	if !taskResponse.CreatedAt.Equal(task.CreatedAt) {
		t.Errorf("Expected CreatedAt to be %v, but got %v", task.CreatedAt, taskResponse.CreatedAt)
	}
	if !taskResponse.UpdatedAt.Equal(task.UpdatedAt) {
		t.Errorf("Expected UpdatedAt to be %v, but got %v", task.UpdatedAt, taskResponse.UpdatedAt)
	}
}

func TestMapTasksToTaskResponse(t *testing.T) {
	tasks := []models.Task{
		{
			ID:          1,
			UserID:      1,
			Title:       "Task 1",
			Description: "Description 1",
			Completed:   false,
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}, {
			ID:          2,
			UserID:      3,
			Title:       "Task 2",
			Description: "Description 2",
			Completed:   false,
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		},
	}

	tasksResponse := MapTasksToTaskResponses(tasks)

	if len(tasksResponse) != len(tasks) {
		t.Errorf("Expected %d task responses, but got %d", len(tasks), len(tasksResponse))
	}

	for i, task := range tasks {
		var mappedTask = MapTaskToTaskResponse(task)

		if !reflect.DeepEqual(tasksResponse[i], mappedTask) {
			t.Errorf("Expected task response to be %+v, but got %+v", mappedTask, tasksResponse[i])
		}
	}
}
