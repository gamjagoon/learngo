package urlshort

import (
  "net/http"
)

func MapHandler (pathToUrls map[string]string, fallback http.Handler)http.HandlerFunc {
  return func (w http.ResponseWriter, r *http.Request)  {
    
  }
}



