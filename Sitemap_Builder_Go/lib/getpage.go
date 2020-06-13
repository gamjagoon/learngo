package lib

import (
  "net/http"
)

// LinkURL : type of Link
type LinkURL struct {
  URL string
  Next *LinkURL
}

// GetWebpage is return page httpResponse
func GetWebpage(url *string) *http.Response {
  resp, err := http.Get(*url)
  if err != nil {
    panic(err)
  }
  return resp
}
