package main

import (
    "fmt"
    "net/http"
    "os"
)

func audioStreamHandler(w http.ResponseWriter, r *http.Request) {
  audioFilePath := "audio.mp3"
  fileInfo, err := os.Stat(audioFilePath)
  if err != nil {
    http.Error(w, "Could not retrieve audio file info", http.StatusInternalServerError)
    return
  }

  // Set the content-type header for the audio file
  w.Header().Set("Content-Type", "audio/mpeg")
  w.Header().Set("Content-Length", fmt.Sprintf("%d", fileInfo.Size()))

  // Serve the file directly
  http.ServeFile(w, r, audioFilePath)
  fmt.Println("Served audio file directly to client")
}


func main() {
  http.HandleFunc("/", audioStreamHandler)
  fmt.Println("Server starting on :8080 for audio streaming")

  if err := http.ListenAndServe(":8080", nil); err != nil {
    fmt.Println("Error starting server:", err)
  }
}
