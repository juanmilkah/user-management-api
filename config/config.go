package config

import (
    "log"
    "os"
    "strconv"
    "github.com/joho/godotenv"
)

type Config struct {
    Port      string
    JWTSecret string
    Env       string
    DB        struct {
        Host     string
        Port     string
        User     string
        Password string
        DBName   string
    }
    RateLimit float64
}

func LoadConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Printf("Error loading .env file, using defaults: %v", err)
    }

    rateLimit, _ := strconv.ParseFloat(getEnv("RATE_LIMIT", "100"), 64)

    cfg := &Config{
        Port:      getEnv("PORT", "8080"),
        JWTSecret: getEnv("JWT_SECRET", "your-secret-key"),
        Env:       getEnv("ENV", "development"),
        RateLimit: rateLimit,
    }

    cfg.DB.Host = getEnv("DB_HOST", "localhost")
    cfg.DB.Port = getEnv("DB_PORT", "5432")
    cfg.DB.User = getEnv("DB_USER", "postgres")
    cfg.DB.Password = getEnv("DB_PASSWORD", "")
    cfg.DB.DBName = getEnv("DB_NAME", "development_database")

    return cfg
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}
