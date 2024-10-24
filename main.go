package main

import (
  "net/http"
  "log"
  "go-server/handlers"
  "github.com/go-chi/chi/v5"
  "github.com/go-chi/chi/v5/middleware"
)

func main(){
  r := chi.NewRouter()

  /*middleware*/ 
  r.Use(middleware.Logger)/*log info on the terminal*/
  r.Use(middleware.Recoverer)

  /*initialise Handler*/ 
  userHandler := handlers.NewUserHandler()

  /*Routes*/ 
  r.Get("/", func(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Connected succesfully!"))
  })

  r.Route("/users", func(r chi.Router){
    r.Get("/", userHandler.GetUsers)
    r.Post("/", userHandler.CreateUser)
    r.Get("/{id}", userHandler.GetUserById)
  })

  /* start server*/ 
  log.Println("Starting server at port 8080...")

  if err := http.ListenAndServe(":8080", r); err != nil{
    log.Fatal(err)
  }
}
