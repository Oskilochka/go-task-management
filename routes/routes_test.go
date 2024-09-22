package routes

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockRegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("User registered successfully"))
	if err != nil {
		return
	}
}

func MockLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Login successful"))
	if err != nil {
		return
	}
}

func TestSetupRouter(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/register", MockRegisterHandler).Methods("POST")
	r.HandleFunc("/login", MockLoginHandler).Methods("POST")

	req, _ := http.NewRequest("POST", "/register", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "User registered successfully", rr.Body.String())

	req, _ = http.NewRequest("POST", "/login", nil)
	rr = httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, "Login successful", rr.Body.String())
}
