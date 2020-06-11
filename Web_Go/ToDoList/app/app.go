package app

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render

type Todo struct {
  ID int `json:"id"`
  Name string `json:"name"`
  CreatedAt time.Time `json:"created_at"`
  Completed bool `json:"completed"`
}

type Success struct{
  Success bool `json:"success"`
}

var todoMap map[int] *Todo

// MakeHandler is return mux.NewRouter
func MakeHandler() http.Handler {

  todoMap = make(map[int]*Todo)
  rd = render.New()
  r := mux.NewRouter()

  r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
  r.HandleFunc("/todos", addTodoHandler).Methods("POST")
  r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
  r.HandleFunc("/", indexHandler)

  return r
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  id, _ := strconv.Atoi(vars["id"])
  if _, ok := todoMap[id]; ok {
    delete(todoMap,id)
    rd.JSON(w, http.StatusOK, Success{true})
  } else {
    rd.JSON(w, http.StatusOK, Success{false})
  }
}

func addTodoHandler(w http.ResponseWriter, r *http.Request) {
  name := r.FormValue("name")
  id := len(todoMap)+1
  todo := &Todo{id, name, time.Now(), false}
  todoMap[id] = todo
  rd.JSON(w, http.StatusOK, todo)
}

func indexHandler(w http.ResponseWriter, r *http.Request){
  http.Redirect(w,r,"/todo.html",http.StatusTemporaryRedirect)
}

func addTestTdos() {
  todoMap[1] = &Todo{1, "Study Golang", time.Now(), false}
  todoMap[2] = &Todo{2, "Eat dinying", time.Now(), true}
  todoMap[3] = &Todo{3, "Go coffe shop", time.Now(), false}
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
  list := []*Todo{}
  addTestTdos()

  for _, v := range todoMap {
    list = append(list, v)
  }

  rd.JSON(w, http.StatusOK, list)
}

