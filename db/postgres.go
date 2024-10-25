package db

import (
  "fmt"
  "github.com/jmoiron/sqlx"
  _ "github.com/lib/pq"
)

type Config struct{
  Host string
  Port string
  User string
  Password string
  DBName string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error){
  dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
    cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
  
  db, err := sqlx.Connect("postgres", dns)
  if err != nil{
    return nil, fmt.Errorf("Error connecting to database %v", err)
  }

  //test the connection
  if err := db.Ping(); err != nil{
    return nil, fmt.Errorf("Error pinging the database %v", err)
  }

  return db, nil
}

//database schema;; 
const schema = `
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
`

func MigrateDB(db *sqlx.DB)error{
  _, err := db.Exec(schema)
  return err
}
