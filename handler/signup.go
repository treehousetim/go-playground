package handler

import (
    "encoding/json"
    "go-webapp-boilerplate/auth"
    "go-webapp-boilerplate/model"
    "net/http"
)

// SignUp handles user registration and returns a QR code for 2FA.
func SignUp(w http.ResponseWriter, r *http.Request) {
    var req SignUpRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := auth.SignUp(DB, req.Email, req.Password); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    secret, qrCode, err := auth.Generate2FASecret(req.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Save the TOTP secret to the user's record in the database
    var user model.User
    DB.Where("email = ?", req.Email).First(&user)
    user.OTPSecret = secret
    user.OTPEnabled = true
    DB.Save(&user)

    // Set the content type to "image/png" and write the QR code image to the response
    w.Header().Set("Content-Type", "image/png")
    w.Write(qrCode)
}

// SignUpRequest represents the expected JSON structure for sign-up requests.
type SignUpRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}
