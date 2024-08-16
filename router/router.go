package router

import (
    "go-webapp-boilerplate/handler"
    "github.com/gorilla/mux"
    "net/http"
)

// SetupRouter initializes the mux router and returns it.
func SetupRouter() *mux.Router {
    r := mux.NewRouter()

    // Define routes
    r.HandleFunc("/api/health", handler.HealthCheck).Methods(http.MethodGet)
    r.HandleFunc("/api/resource", handler.GetResource).Methods(http.MethodGet)
    r.HandleFunc("/api/resource", handler.CreateResource).Methods(http.MethodPost)
    r.HandleFunc("/api/signup", handler.SignUp).Methods(http.MethodPost)
    r.HandleFunc("/api/signin", handler.SignIn).Methods(http.MethodPost)

    return r
}
