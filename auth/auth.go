package auth

import (
    "go-webapp-boilerplate/model"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "errors"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func SignUp(db *gorm.DB, email, password string) error {
    hashedPassword, err := HashPassword(password)
    if err != nil {
        return err
    }

    user := model.User{
        Email:    email,
        Password: hashedPassword,
    }

    return db.Create(&user).Error
}

func SignIn(db *gorm.DB, email, password string) (string, error) {
    var user model.User
    err := db.Where("email = ?", email).First(&user).Error
    if err != nil {
        return "", err
    }

    if !CheckPasswordHash(password, user.Password) {
        return "", errors.New("invalid credentials")
    }

    return GenerateJWT(user.Email)
}
