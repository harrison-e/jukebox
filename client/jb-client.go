package main 

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

func main() {
  resp, err := http.Get("http://localhost:8080/Go")
  if err != nil {
    fmt.Println("HTTP Get Error:", err)
    return
  }
  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("Error reading body:", err)
    return
  }

  fmt.Println("Server says:", string(body))
}
