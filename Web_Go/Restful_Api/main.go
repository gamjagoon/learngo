package main

import (
	"net/http"
	"Restful_Api/src"
)

func main()  {
	http.ListenAndServe(":8080",myapp.NewHandler())
}