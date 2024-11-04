package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
  http.HandleFunc("/", handler)
  fmt.Println("Server starting on :8080")

  if err := http.ListenAndServe(":8080", nil); err != nil {
    fmt.Println("Error starting server:", err)
  }
}
