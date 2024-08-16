package main

import (
    "go-webapp-boilerplate/handler"
    "go-webapp-boilerplate/middleware"
    "go-webapp-boilerplate/model"
    "go-webapp-boilerplate/router"
    "log"
    "net/http"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    // Setup Database Connection using SQLite
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Automigrate the database schema
    db.AutoMigrate(&model.User{}, &model.Resource{})

    // Set the global DB connection for handlers
    handler.DB = db

    // Setup Router
    r := router.SetupRouter()

    // Setup CORS Middleware
    handler := middleware.SetupCORS(r)

    // Start the server
    log.Println("Server running on port 8080")
    if err := http.ListenAndServe(":8080", handler); err != nil {
        log.Fatal(err)
    }
}
