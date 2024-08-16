package handler

import (
    "encoding/json"
    "net/http"
    "go-webapp-boilerplate/auth"
    "go-webapp-boilerplate/model"
)


// SignInRequest represents the expected JSON structure for sign-in requests.
type SignInRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
    OTPCode  string `json:"otp_code"` // Optional: Only required if 2FA is enabled
}

// SignIn handles user login, including password and 2FA verification.
func SignIn(w http.ResponseWriter, r *http.Request) {
    var req SignInRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user model.User
    if err := DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        http.Error(w, "user not found", http.StatusUnauthorized)
        return
    }

    // Verify the password
    if !auth.CheckPasswordHash(req.Password, user.Password) {
        http.Error(w, "invalid credentials", http.StatusUnauthorized)
        return
    }

    // If 2FA is enabled, verify the OTP code
    if user.OTPEnabled {
        if !auth.Validate2FACode(user.OTPSecret, req.OTPCode) {
            http.Error(w, "invalid OTP code", http.StatusUnauthorized)
            return
        }
    }

    // Generate a JWT token (or session token)
    token, err := auth.GenerateJWT(user.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return the token as a JSON response
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
