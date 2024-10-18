package handlers

import (
	"backend/config"
	"backend/models"
	"backend/repository"
	"backend/utils"
	"encoding/json"
	"fmt"
	"log"
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

	// Generate authorization token
	authToken := utils.GenerateResetToken()
	expirationTime := time.Now().Add(24 * time.Hour) // Token valid for 24 hours

	user.AuthToken = authToken
	user.AuthTokenExpiration = expirationTime
	user.IsVerified = false // Mark user as unverified initially

	if err := repository.CreateUser(user); err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		return
	}

	// Send email with verification link
	verificationLink := fmt.Sprintf("http://localhost:3000/verify-email?token=%s", authToken)
	emailBody := fmt.Sprintf(`
        <h2>Email Verification</h2>
        <p>Click the link below to verify your email:</p>
        <a href="%s">Verify Email</a>
        <p>This link will expire in 24 hours.</p>
    `, verificationLink)

	if err := config.SendEmail(user.Email, "Email Verification", emailBody); err != nil {
		log.Printf("Error sending email: %v", err)
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully. Please check your email to verify your account."})
}

func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	userID, err := repository.ValidateAuthToken(token)
	if err != nil {
		http.Error(w, "Invalid or expired token", http.StatusBadRequest)
		return
	}

	if err := repository.VerifyUser(userID); err != nil {
		http.Error(w, "Error verifying user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Email verified successfully"})
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

	resetLink := fmt.Sprintf("http://localhost:3000/reset-password?token=%s", resetToken)
	emailBody := fmt.Sprintf(`
		<h2>Password Reset Request</h2>
		<p>Click the link below to reset your password:</p>
		<a href="%s">Reset Password</a>
		<p>This link will expire in 15 minutes.</p>
		<p>If you didn't request this, please ignore this email.</p>
	`, resetLink)

	if err := config.SendEmail(user.Email, "Password Reset Request", emailBody); err != nil {
		log.Printf("Error sending email: %v", err)
		http.Error(w, "Error sending email", http.StatusInternalServerError)
		return
	}

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
