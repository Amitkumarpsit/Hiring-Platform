package handlers

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"encoding/json"
	"net/http"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := user.HashPassword(); err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	if err := repository.CreateUser(user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		LoginID  string `json:"loginID"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByLoginID(credentials.LoginID)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := user.ComparePassword(credentials.Password); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.ID.Hex())
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := repository.GetUserByEmail(request.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	resetToken := utils.GenerateResetToken()
	expirationTime := time.Now().Add(15 * time.Minute)

	if err := repository.SaveResetToken(user.ID, resetToken, expirationTime); err != nil {
		http.Error(w, "Error saving reset token", http.StatusInternalServerError)
		return
	}

	// TODO: Send email with reset token (implement email sending functionality)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password reset instructions sent to email"})
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ResetToken  string `json:"resetToken"`
		NewPassword string `json:"newPassword"`
	}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, err := repository.ValidateResetToken(request.ResetToken)
	if err != nil {
		http.Error(w, "Invalid or expired reset token", http.StatusBadRequest)
		return
	}

	if err := repository.UpdatePassword(userID, request.NewPassword); err != nil {
		http.Error(w, "Error updating password", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Password reset successfully"})
}
