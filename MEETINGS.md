# Meeting Notes 


## 12 November 2024
**Discussed:**
- Beep basics
- WebSocket basics
- Client/server interactions

**Basic client join:**
1. Client connects to server
2. Server accepts client connection, sends initial timestamp and first TWO audio files
3. Client notes timestamp, and begins downloading files
4. Client finishes downloading files, and seeks audio playback stream (not online stream, beep stream) to timestamp + time taken to download
5. If timestamp + time taken to download is greater than first song length, begin playback of second song at 0 + remainder
6. If timestamp + time taken to download is greater than first song length + second song length, abort and reinit
7. Client plays song, begins downloading rest of queue in background

**Goals for 17 November:**
- Serve mp3 over WebSocket
- Server-side timestamp and queue
- Pause/play


## 3 November 2024
**Discussed:**
- Initial roadmapping

**Goals for 10 November:**
- Comfortable with basic Go programs, ideally using sockets 
- Have research prepared for audio streaming options
- Have an idea of how to structure a Go project like ours (client and server in 1 repo)
