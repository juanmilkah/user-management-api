package models

type User struct{
  Id string `json:"id"`
  Name string `json:"name" validate:"required,min=2,max=50"`
  Email string `json:"string" validate:required,email"`
  Password string `json:"password,omitempty" validate:"required,min=6"`
}

type LoginRequest struct{
  Email string `json:email validate:"required,email"`
  Password string `json:password validate:required`
}

type TokenResponse struct {
  Token string `json:token`
}

