package main

import (
    "fmt"
    "log"
    "net/http"
    "go-server/config"
    "go-server/db"
    "go-server/handlers"
    "go-server/middleware"
    "go-server/models"
    "go-server/utils"
    "github.com/go-chi/chi/v5"
    chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Connected successfully\n")
}

func main() {
    // Load configuration
    cfg := config.LoadConfig()

    // Initialize database
    dbConfig := db.Config{
        Host:     cfg.DB.Host,
        Port:     cfg.DB.Port,
        User:     cfg.DB.User,
        Password: cfg.DB.Password,
        DBName:   cfg.DB.DBName,
    }

    database, err := db.NewPostgresDB(dbConfig)
    if err != nil {
        log.Fatal(err)
    }
    defer database.Close()

    // Run migrations
    if err := db.MigrateDB(database); err != nil {
        log.Fatal(err)
    }

    // Initialize repositories
    userRepo := models.NewUserRepository(database)

    // Initialize error handler
    errorHandler := utils.NewErrorHandler(cfg.Env)

    // Initialize router
    r := chi.NewRouter()

    // Initialize rate limiter
    rateLimiter := middleware.NewRateLimiter(cfg.RateLimit, errorHandler)

    // Global middleware
    r.Use(chimiddleware.Logger)
    r.Use(chimiddleware.Recoverer)
    r.Use(chimiddleware.RequestID)
    r.Use(rateLimiter.RateLimit)

    // Initialize handlers and middleware
    userHandler := handlers.NewUserHandler(cfg.JWTSecret, errorHandler, userRepo)
    authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret, errorHandler)

    // Public routes
    r.Group(func(r chi.Router) {
        r.Get("/", homeHandler)
        r.Post("/login", userHandler.Login)
        r.Post("/users/new", userHandler.CreateUser)
    })

    // Protected routes
    r.Group(func(r chi.Router) {
        r.Use(authMiddleware.RequireAuth)
        
        r.Route("/users", func(r chi.Router) {
            r.Get("/", userHandler.GetUsers)
            r.Get("/{id}", userHandler.GetUserById)
        })
    })

    // Start server
    addr := fmt.Sprintf(":%s", cfg.Port)
    log.Printf("Server starting on port %s...", cfg.Port)
    if err := http.ListenAndServe(addr, r); err != nil {
        log.Fatal(err)
    }
}
