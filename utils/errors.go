package utils

import (
  "encoding/json"
  "net/http"
)

type ErrorResponse struct{
  Status int `json:"status"`
  Message string `"json:message"`
  Error string `json:"error,omitempty"`
}

type ErrorHandler struct{
  env string
}

func NewErrorHandler(env string) *ErrorHandler{
  return &ErrorHandler{
    env : env,
  }
}

func(h *ErrorHandler) RespondWithError(w http.ResponseWriter, status int, message string, err error){
  resp := ErrorResponse{
    Status: status,
    Message: message,
  }

  if err != nil && resp.Status == http.StatusInternalServerError{
    if h.env == "development" {
      resp.Error = err.Error()
    }else{
      resp.Error = "Internal Server Error"
    }
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  json.NewEncoder(w).Encode(resp)
}

func(h *ErrorHandler) RespondWithJSON(w http.ResponseWriter, status int, payload interface{}){
  w.Header().Set("Content-TYpe", "application/json")
  w.WriteHeader(status)
  json.NewEncoder(w).Encode(payload)
}
