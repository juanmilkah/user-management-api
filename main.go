package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
)

/*define User*/
type User struct {
  Id string `json:id`
  Name string `json:name`
  Email string `josn:email`
}

func homeHandler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Welcome to my server");
}

func getUsers(w http.ResponseWriter, r *http.Request){
  users := []User{
    {Id: "1", Name: "mike", Email: "mike@gmail.com"},
    {Id: "2", Name: "nancy", Email: "nancy@gmail.com"},
  }

  w.Header().Set("Content-Type","application/json");

  /*encode and send json data*/ 
  json.NewEncoder(w).Encode(users);
}

func createUser(w http.ResponseWriter, r *http.Request){
  /*allow POST method only*/ 
  if r.Method != http.MethodPost{
    http.Error(w, "Method not allowed", http.StatusMethodNotAllowed);
    return;
  }

  var newUser User;
  /*parse the body*/
  if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil{
    http.Error(w, err.Error(), http.StatusBadRequest);
    return;
  }

  //add to database

  w.Header().Set("Content-Type", "application/json");
  w.WriteHeader(http.StatusCreated);

  /*send user back*/ 
  json.NewEncoder(w).Encode(newUser);

}

func main(){
  /*define routes*/
  http.HandleFunc("/", homeHandler);
  http.HandleFunc("/users", getUsers);
  http.HandleFunc("/users/create", createUser);

  //start the server at port 8000
  log.Println("Server started at port 8000")
  if err := http.ListenAndServe(":8000", nil); err != nil {
    log.Fatal(err);
  }
}
