package handlers

import (
  "net/http"
  "encoding/json"
  "go-server/models"
  "github.com/go-chi/chi/v5"
)

type UserHandler struct {}

func NewUserHandler() *UserHandler{
  return &UserHandler{};
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request){
  users := []models.User{
     {Id: "1", Name: "John Doe", Email: "john@example.com"},
     {Id: "2", Name: "Jane Doe", Email: "jane@example.com"},
  }

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request){
  var newUser models.User;

  if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil{
    http.Error(w, err.Error(), http.StatusBadRequest);
    return
  }

  /* add to database or something*/ 
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)

  json.NewEncoder(w).Encode(newUser)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request){
  /* get Id from url*/ 
  userId := chi.URLParam(r, "id")

    /*fetch frrom database*/ 
    user:= models.User{
      Id: userId,
      Name: "Nadreas",
      Email: "Nadreas@me.com",
    }

    w.Header().Set("Content-Type" , "application/json")
    json.NewEncoder(w).Encode(user)
}
