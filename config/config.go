package config

import (
  "log"
  "os"
  "github.com/joho/godotenv"
)

type Config struct{
  Port  string
  JwtSecret string
  Env string
}

func LoadConfig() *Config{
  err := godotenv.Load()

  if err != nil{
    log.Printf("Error loading .env file: using defaults: %v", err)
  }

  return &Config{
    Port: getEnv("PORT", "8080"),
    JwtSecret: getEnv("JWT_SECRET", "secret_password"),
    Env: getEnv("ENV", "development"),
  }
}

func getEnv(key, defaultValue string) string {
  value := os.Getenv(key)

  if value == "" {
    return defaultValue
  }
  return value 
}

