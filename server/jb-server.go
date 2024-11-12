package main

import (
    "fmt"
    "os"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader {
  CheckOrigin: func(r *http.Request) bool {
    return true // Allow connections from all origins (for now)
  },
}

func echoWebSocketHandler(w http.ResponseWriter, r *http.Request) {
  // Upgrade HTTP to WebSocket
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    fmt.Println("Error upgrading connection:", err)
    return
  }
  defer conn.Close()
  fmt.Println("Established new WebSocket connection")

  // Infinite read/send loop (as WS stays open)
  for {
    // ReadMessage blocks until receiving msg
    messageType, msg, err := conn.ReadMessage()
    if err != nil {
      fmt.Println("Error reading message from client:", err)
      break
    }

    // Log client message
    fmt.Printf("Received: \"%s\"\n", msg)
    
    // Echo message back to client 
    err = conn.WriteMessage(messageType, msg)
    if err != nil {
      fmt.Println("Error writing message to client:", err)
      break
    }
  }
}

func audioStreamRawHttpHandler(w http.ResponseWriter, r *http.Request) {
  // Stat audio.mp3
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
  http.HandleFunc("/", audioStreamRawHttpHandler)
  http.HandleFunc("/ws", echoWebSocketHandler)
  fmt.Println("Server starting on :8080")

  if err := http.ListenAndServe(":8080", nil); err != nil {
    fmt.Println("Error starting server:", err)
  }
}
