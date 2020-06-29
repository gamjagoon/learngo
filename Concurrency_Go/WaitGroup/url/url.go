package url

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func sendRequest(url string) {
  defer wg.Done()
  res, err := http.Get(url)
  if err != nil {
    panic(err)
  }

  fmt.Printf("[%d] %s\n",res.StatusCode, url)
}

func main() {
  
  urls := [5]string{"google.com","youtube.com","naver.com","daum.net","twitter.com" }
  for _, url := range urls{

    // // not gorutine
    //sendRequest("https://" + url)
    
    // gorutine
    go sendRequest("https://" + url)
    wg.Add(1)
  }

  wg.Wait()
}