package handlers

import (
  "net/http"
  "encoding/json"
  "strconv"
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
  userRepo *models.UserRepository
}

func NewUserHandler(jwtSecret string, errorHandler *utils.ErrorHandler, userRepo *models.UserRepository) *UserHandler{
  return &UserHandler{
    validate: validator.New(),
    jwtSecret: jwtSecret,
    errorHandler: errorHandler,
    userRepo: userRepo,
  };
}

//get users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request){
  users, err := h.userRepo.GetAll()
  
  if err != nil{
    h.errorHandler.RespondWithError(w, http.StatusInternalServerError, "Could not fetch users", err)
    return 
  }

  h.errorHandler.RespondWithJSON(w, http.StatusOK, users)
}

// Update CreateUser method
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        h.errorHandler.RespondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    if err := h.validate.Struct(user); err != nil {
        h.errorHandler.RespondWithError(w, http.StatusBadRequest, "Validation failed", err)
        return
    }

    if err := h.userRepo.Create(user); err != nil {
        h.errorHandler.RespondWithError(w, http.StatusInternalServerError, "Could not create user", err)
        return
    }

    h.errorHandler.RespondWithJSON(w, http.StatusCreated, user)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request){
  /* get Id from url*/ 
  userId := chi.URLParam(r, "id")
  id, err := strconv.Atoi(userId)

  if err != nil{
    panic(err)
  }

  /*fetch frrom database*/ 
   user, err := h.userRepo.GetById(id) 
  if err != nil{
    h.errorHandler.RespondWithError(w, http.StatusInternalServerError, "Could not fetch User", err)
    return 
  }

  h.errorHandler.RespondWithJSON(w, http.StatusOK, user)
}

// Update Login method
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
    var login models.LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
        h.errorHandler.RespondWithError(w, http.StatusBadRequest, "Invalid request payload", err)
        return
    }

    if err := h.userRepo.VerifyPassword(login.Email, login.Password); err != nil {
        h.errorHandler.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials", nil)
        return
    }

    user, err := h.userRepo.GetByEmail(login.Email)
    if err != nil {
        h.errorHandler.RespondWithError(w, http.StatusInternalServerError, "Could not get user", err)
        return
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.Id,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString([]byte(h.jwtSecret))
    if err != nil {
        h.errorHandler.RespondWithError(w, http.StatusInternalServerError, "Could not generate token", err)
        return
    }

    h.errorHandler.RespondWithJSON(w, http.StatusOK, models.TokenResponse{Token: tokenString})
}
