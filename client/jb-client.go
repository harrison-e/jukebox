package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
  "time"
  "github.com/gopxl/beep"
  "github.com/gopxl/beep/mp3"
  "github.com/gopxl/beep/speaker"
)

func main() {
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

  // Move file cursor back to start 
  tempFile.Seek(0, io.SeekStart)

  // Decode mp3 into streamer with beep/mp3
  streamer, format, err := mp3.Decode(tempFile)
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
