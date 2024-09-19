package handlers

import (
	"encoding/json"
	"josk/task-management-system/auth"
	"josk/task-management-system/database"
	"josk/task-management-system/models"
	"josk/task-management-system/utils"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid JSON format"}, http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Error hashing password"}, http.StatusInternalServerError)
		return
	}

	user.Password = hashedPassword

	result := database.DB.Create(&user)

	if result.Error != nil {
		utils.SendJSONResponse(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, map[string]string{"message": "User was created"}, http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.User

	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid JSON format"}, http.StatusBadRequest)
		return
	}

	var user models.User
	database.DB.Where("username = ?", credentials.Username).First(&user)

	if user.ID == 0 {
		utils.SendJSONResponse(w, map[string]string{"error": "User was not found"}, http.StatusNotFound)
		return
	}

	if !utils.CheckPassHash(credentials.Password, user.Password) {
		utils.SendJSONResponse(w, map[string]string{"error": "Invalid email or password"}, http.StatusUnauthorized)
		return
	}

	tokenStr, err := auth.GenerateJWT(user.ID, user.Username)

	if err != nil {
		utils.SendJSONResponse(w, map[string]string{"error": "Token was not created"}, http.StatusInternalServerError)
		return
	}

	utils.SendJSONResponse(w, map[string]string{"token": tokenStr}, http.StatusOK)
}
