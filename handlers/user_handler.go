package handlers

import (
  "net/http"
  "encoding/json"
  "go-server/models"
  "go-server/utils"
  "time"
  "github.com/golang-jwt/jwt/v5"
  "github.com/go-playground/validator/v10"
  "github.com/go-chi/chi/v5"
)

type UserHandler struct {
  validate *validator.Validate
  jwtSecret string
  errorHandler *utils.ErrorHandler
}

func NewUserHandler(jwtSecret string, errorHandler *utils.ErrorHandler) *UserHandler{
  return &UserHandler{
    validate: validator.New(),
    jwtSecret: jwtSecret,
    errorHandler: errorHandler,
  };
}

//get users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request){
  users := []models.User{
     {Id: "1", Name: "John Doe", Email: "john@example.com"},
     {Id: "2", Name: "Jane Doe", Email: "jane@example.com"},
  }

  w.Header().Set("Content-Type", "application/json")
  h.errorHandler.RespondWithJSON(w, http.StatusOK, users)
}

//crerate a user
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request){
  var newUser models.User;

  if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil{
    h.errorHandler.RespondWithError(w, http.StatusBadRequest, "Invalid Payload", err)
    return
  }

  //validate Payload
  if err := h.validate.Struct(newUser); err != nil{
    h.errorHandler.RespondWithError(w, http.StatusBadRequest, "Validation failed", err)
    return 
  }

  /* add password to database or something*/ 

  //respond 
  h.errorHandler.RespondWithJSON(w, http.StatusCreated, newUser)
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
    h.errorHandler.RespondWithJSON(w, http.StatusOK, user)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request){
  var login models.LoginRequest

  if err := json.NewDecoder(r.Body).Decode(&login); err != nil{
    h.errorHandler.RespondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
    return 
  }

  //validate data
  if err := h.validate.Struct(login); err != nil{
    h.errorHandler.RespondWithError(w, http.StatusBadRequest, "Validation failed", err)
    return
  }
  
  //verify password against database

  //create jwt token
  token := jwt.NewWithClaims(jwt.SigningMethodHS256,  jwt.MapClaims{
    "user_id": "123" ,
    "exp": time.Now().Add(time.Hour * 24).Unix(),
  })

  tokenString, err := token.SignedString([]byte(h.jwtSecret))
  if err != nil{
    h.errorHandler.RespondWithError(w, http.StatusInternalServerError, "Could not generate Token", err)
    return 
  }

  h.errorHandler.RespondWithJSON(w, http.StatusOK, models.TokenResponse{Token: tokenString})
}
