package models

import (
  "database/sql"
  "errors"
  "time"
  "github.com/jmoiron/sqlx"
  "golang.org/x/crypto/bcrypt"
  )

type User struct{
  Id string `json:"id"`
  Name string `json:"name" validate:"required,min=2,max=50"`
  Email string `json:"email" validate:"required,email"`
  Password string `json:"password,omitempty" validate:"required,min=6"`
  PasswordHash string `json:"-" db:"password_hash"`
  CreatedAt time.Time `json:"created_at" db:"created_at"`
}
 
type UserRepository struct{
  db *sqlx.DB 
}

func NewUserRepository(db *sqlx.DB) *UserRepository{
  return &UserRepository{db:db }
}

func(r *UserRepository) Create(user User) error{
  //hash password
  hashed_password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err != nil{
    return err
  }

   query := `
        INSERT INTO users (name, email, password_hash)
        VALUES ($1, $2, $3)
        RETURNING id, created_at`

  return r.db.QueryRow(
    query,
    user.Name,
    user.Email,
    string(hashed_password),
  ).Scan(&user.Id, &user.CreatedAt)
}

func (r *UserRepository) GetById(id int) (*User, error){
  var user User
  err := r.db.Get(&user,
      "SELECT id, name, email, created_at FROM users WHERE id = $1",
      id)
  if err == sql.ErrNoRows{
    return nil, errors.New("User not found")
  }

  return &user, nil
}

func (r *UserRepository) GetByEmail(email string) (*User, error){
  var user User

  err := r.db.Get(&user,
      "SELECT id, name, email, password_hash, created_at FROM users WHERE email = $1",
      email)
  if err == sql.ErrNoRows {
      return nil, errors.New("user not found")
    }
  return &user, err
}

func (r *UserRepository) GetAll() ([]User, error) {
    var users []User
    err := r.db.Select(&users,
        "SELECT id, name, email, created_at FROM users ORDER BY created_at DESC")
    return users, err
}

func (r *UserRepository) VerifyPassword(email, password string) error {
  var hash string
  err := r.db.Get(&hash,
     "SELECT password_hash FROM users WHERE email = $1",
      email)
  if err != nil {
    return err
  }

  return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

type LoginRequest struct{
  Email string `json:"email" validate:"required,email"`
  Password string `json:"password" validate:"required"`
}

type TokenResponse struct {
  Token string `json:"token"`
}

