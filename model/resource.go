package model

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email       string `json:"email" gorm:"unique"`
    Password    string `json:"-"`
    OTPSecret   string `json:"-"`
    OTPEnabled  bool   `json:"otp_enabled"`
}

type Resource struct {
    gorm.Model
    Name        string `json:"name"`
    Description string `json:"description"`
}
