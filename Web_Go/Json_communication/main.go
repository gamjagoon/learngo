package main

import (
	"Json_communication/myapp"
	"net/http"
)

const port string = ":10001"

func main() {

	http.ListenAndServe(port, myapp.NewHTTPHandler())
}
