package main

/*
text/template => example.tmpl copy and write
html/template => write parse case by html tags 
*/
import (
	"html/template"
	"os"
)

type User struct {
  Name string
  Email string
  Age int
}

func (u User) IsOld() bool {
  return u.Age > 30
}

var( 
  use = User{
    Name:"Minjae",
    Email:"rlaalswo201@gmail.com",
    Age:23,
  }
  use2 = User{
    Name:"DaeWon",
    Email:"dawon@gmail.com",
    Age:40,
  }
)

func main() {
  tmpl, err := template.New("Tmp1").ParseFiles("./tmplate/tmp1.tmpl","./tmplate/tmp2.tmpl","./tmplate/tmp3.tmpl")
  users := []User{use, use2}
  if err != nil {
    panic(err)
  }
  tmpl.ExecuteTemplate(os.Stdout,"tmp2.tmpl",use)
  tmpl.ExecuteTemplate(os.Stdout,"tmp2.tmpl",use2)
  // List data push
  tmpl.ExecuteTemplate(os.Stdout,"tmp3.tmpl",users)
}
