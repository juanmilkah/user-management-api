package main

import (
  "fmt"
  "net/http"
  "log"
  "go-server/handlers"
  "go-server/middleware"
  "go-server/config"
  "go-server/utils"
  "github.com/go-chi/chi/v5"
  chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func main(){
  //load config
  cfg := config.LoadConfig()
  //initialize router
  r := chi.NewRouter()

  //error handler
  errorHandler := utils.NewErrorHandler(cfg.Env)

  /*middleware*/ 
  r.Use(chimiddleware.Logger)/*log info on the terminal*/
  r.Use(chimiddleware.Recoverer)
  r.Use(chimiddleware.RequestID)

  /*initialise Handler*/ 
  userHandler := handlers.NewUserHandler(cfg.JwtSecret, errorHandler)
  authMiddleware := middleware.NewAuthMiddleware(cfg.JwtSecret, errorHandler)

  /*Public Routes*/
  r.Group(func(r chi.Router){
    r.Post("/login", userHandler.Login)
    r.Post("/users/create", userHandler.GetUsers)
  })

  /*protected routes*/
  r.Group(func(r chi.Router){
    r.Use(authMiddleware.RequireAuth)

    r.Route("/users", func(r chi.Router){
      r.Get("/", userHandler.GetUsers)
      r.Get("/{id}", userHandler.GetUserById)
    })
  })

  /*Routes*/ 
  r.Get("/", func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Connected succesfully!\n"))
  })

  /* start server*/ 
  addr := fmt.Sprintf(":%s", cfg.Port)
  log.Printf("Starting server at port %v ...", cfg.Port)

  if err := http.ListenAndServe(addr, r); err != nil{
    log.Fatal(err)
  }
}
