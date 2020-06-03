package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
)

var rd *render.Render

type User struct {
  Name string `json:"name"`
  Email string `json:"email"`
  CreatedAt time.Time `json:"created_at`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request)  {
  user := User{Name:"minjae", Email:"rlaalswo201@gmail.com"}

  // w.Header().Add("Content-type", "apllication/json")
  // w.WriteHeader(http.StatusOK)
  // data, _ := json.Marshal(user)
  // fmt.Fprint(w, string(data))

  // // repalce 1 line !! render

  rd.JSON(w, http.StatusOK, user)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
  user := new(User)
  err := json.NewDecoder(r.Body).Decode(user)
  if err != nil {
    rd.Text(w, http.StatusBadRequest, err.Error())
    return
  }
  user.CreatedAt = time.Now()
  rd.JSON(w, http.StatusOK, user)
}

func helloHandler (w http.ResponseWriter, r *http.Request) {
  user := User{Name: "Minjae", Email: "rlaalswo201@gmail.com"}
  // tmpl, err := template.New("Hello").ParseFiles("./templates/hello.tmpl")
  // if err != nil{
  //   rd.Text(w, http.StatusBadRequest, err.Error())
  //   return
  // }
  // tmpl.ExecuteTemplate(w, "hello.tmpl", "Minjae")
  rd.HTML(w, http.StatusOK, "body", user)
}

func main() {
  rd = render.New(render.Options{
    Directory: "templates",
    Extensions: []string{".html"},
    Layout: "hello",
  })
  mux := pat.New()

  mux.Get("/users", getUserInfoHandler)
  mux.Post("/users",addUserHandler)
  mux.Get("/hello", helloHandler)

  http.ListenAndServe(":8080", mux)
}