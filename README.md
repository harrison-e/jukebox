# Jukebox
Jukebox is an open-source, self-hosted music streaming server model.

## Proposed tech stack
Two clients:
1. Webapp client in frontend JavaScript/HTML/CSS
2. TUI client in whatever lang (Go or Rust)

Backend:
- Go backend server logic
- Relational database for storage/querying (MySQL? PostgreSQL?)

Use WebRTC, RTP, or some other streaming paradigm for client/server sockets

```

              Host server                   Connected clients
    +-----------------------------+
    |   +---------------------+   |
    |   |  postgres music db  |   |           +----------+
    |   +---------------------+   |       +-->| jbclient |
    |     | | |         A A A     |      /    +----------+
    |     | | |         | | |     |     / 
    |     V V V         | | |     |    /
    |   +---------------------+   |   /       +----------+
    |   |       jbserver      |<-----+------->| jbclient |
    |   +---------------------+   |   \       +----------+
    +-----------------------------+    \
                                        \ 
                                         \    +----------+
                                          +-->| jbclient |
                                              +----------+

```

## Roadmap                                    
1. **Basic terminal client/server in Go**
    - Send HTTP packets back and forth
    - How to use Go? WTF is Go? 
    - How to structure Go Project? We want client and server in the same repo.
2. **Audio streaming over TCP/IP/HTTP from server to client**
    - Figure out what streaming paradigm to use. WebRTC? RTP? Raw HTTP?
    - Server will have a few songs to choose from 
    - Client can choose which song to play, toggle playback
3. **Working music queue** 
    - Server uses a queue of songs to stream to connected clients
    - Clients can queue songs on server, affecting playback for all connected and future clients
4. **Port forward**
    - Open up Jukebox server to the Internet
    - Look into security concerns
5. **Multiple sessions on one server** 
    - User joins server, can choose to start a new session or join any existing ones
6. **Implement a relational database**
    - Implement server-side interactions with an existing relational database 
    - Used for querying the database (searching by artist, date, genre, etc.) 
7. **Implement pretty frontend**
    - Mixed JavaScript frontend application to serve the same purpose as the terminal client
    - Spotify-esque UI? Need to look into/develop this
8. **Implement playlists and nested playlists**
    - Implement classic playlists (song lists) as well as nested playlists 
    - *Classic playlists:* a list of songs 
    - *Nested playlists:* entries can be songs or playlists 
    - Give users the ability to squash/merge nested playlists (like Photoshop layers)
    - Treat playlists as building blocks
    - *Implementation ideas:* JSON, Filesystem directory analogy 

## Inspiration
- [Jukebox.today (unmaintained peaked-in-high-school-ass version of this project)](http://jukebox.today)
- [Groove Basin (old version of what we want to do)](https://github.com/andrewrk/groovebasin)
- [Nate's Zoom project](./res/proj3_report.pdf)
