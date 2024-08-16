package handler

import (
    "encoding/json"
    "net/http"
    "go-webapp-boilerplate/model"
    "gorm.io/gorm"
)

var DB *gorm.DB

// HealthCheck handler returns a simple health check status.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
    response := map[string]string{"status": "OK"}
    json.NewEncoder(w).Encode(response)
}

// GetResource retrieves a resource.
func GetResource(w http.ResponseWriter, r *http.Request) {
    var resources []model.Resource
    DB.Find(&resources)
    json.NewEncoder(w).Encode(resources)
}

// CreateResource creates a new resource.
func CreateResource(w http.ResponseWriter, r *http.Request) {
    var resource model.Resource
    if err := json.NewDecoder(r.Body).Decode(&resource); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    DB.Create(&resource)
    json.NewEncoder(w).Encode(resource)
}


