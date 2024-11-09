package main

import (
  "fmt"
  "io"
  "os"
  "time"
  "net/http"

  "github.com/gopxl/beep"
  "github.com/gopxl/beep/mp3"
  "github.com/gopxl/beep/speaker"

  "github.com/gorilla/websocket"
)

func playMp3File(f *os.File) {
  // Move file cursor back to start 
  f.Seek(0, io.SeekStart)

  // Decode mp3 into streamer with beep/mp3
  streamer, format, err := mp3.Decode(f)
  if err != nil {
    fmt.Println("Error decoding mp3:", err)
    return
  }
  defer streamer.Close()

  // Initialize beep/speaker
  // First arg is sample rate, second is buffer size
  speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

  // Make a channel to be notified when Seq is over 
  // - beep.Seq takes a sequence of streamers and plays them in order 
  // - beep.Callback creates a "streamer" that is really a callback fn
  //   ^ this might be super useful for client/server stuff
  done := make(chan bool)
  speaker.Play(beep.Seq(streamer, beep.Callback(func() {
    done <- true
  })))

  <-done
}

func retrieveAudioRawHttp() {
  resp, err := http.Get("http://localhost:8080/") // Connect to the server stream
  if err != nil {
    fmt.Println("Error connecting to server:", err)
    return
  }
  defer resp.Body.Close()

  tempFile, err := os.CreateTemp("", "streamed_audio_*.mp3")
  if err != nil {
    fmt.Println("Error creating temp file:", err)
    return
  }
  defer os.Remove(tempFile.Name()) // Clean up temp file
  defer tempFile.Close()

  // Write the full stream to the temp file
  _, err = io.Copy(tempFile, resp.Body)
  if err != nil {
    fmt.Println("Error streaming audio:", err)
    return
  }

  // Play temp file 
  playMp3File(tempFile)
}

func connectEchoWebSocket() {
  // Connect to WS server
  conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
  if err != nil {
    fmt.Println("WS dial failed:", err)
    return
  }
  defer conn.Close()

  // Send message to server
  err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, WS!"))
  if err != nil {
    fmt.Println("Failed to write to WS:", err)
    return
  }
  fmt.Println("Sent message to server")

  // Set read timeout
  conn.SetReadDeadline(time.Now().Add(15 * time.Second))

  // Read echo from server
  _, msg, err := conn.ReadMessage()
  if err != nil {
    fmt.Println("Failed to read from WS:", err)
    return
  }
  fmt.Printf("Received message from server: \"%s\"\n", msg)

  // Keep connection open
  for {}
}

func main() {
  connectEchoWebSocket()
}
