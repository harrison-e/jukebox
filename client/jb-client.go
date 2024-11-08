package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "os/exec"
    "runtime"
)

func main() {
    resp, err := http.Get("http://localhost:8080/Go") // Connect to the server stream
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

    // Ensure the temp file is fully written before playing
    tempFile.Close()

    // Play the audio file based on OS
    var cmd *exec.Cmd
    if runtime.GOOS == "darwin" {
        cmd = exec.Command("afplay", tempFile.Name())
    } else if runtime.GOOS == "linux" {
        cmd = exec.Command("cvlc", tempFile.Name())
    } else {
        fmt.Println("Unsupported OS")
        return
    }

    // Run the playback command
    err = cmd.Run()
    if err != nil {
        fmt.Println("Failed to play audio:", err)
    }
}
