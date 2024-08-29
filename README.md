# FreeTube Sync

[FreeTube](https://freetubeapp.io/) is an open source desktop YouTube player built with privacy in mind. Use YouTube without advertisements and prevent Google from tracking you with their cookies and JavaScript. 

It does not require an account, all datas are on your local storage. In case of you use FreeTube on several computers, you can't synchronize them easily. FreeTube stores datas on plain text files, loaded in memory and rewrite them on each updates.

**FreeTube Sync** tries to solve this problem.

## What are requirements?

**FreeTube Sync** requires a server accessible by all machines running FreeTube (on a local network, through a VPN or on the web).

## How does it work?

The role of the server is to store the history, the playlists and the profiles of FreeTube instances (clients).

After starting the server, each client must init its local datas on the server. This action must be processed only once.
At each time you want to use FreeTube, you have to pull datas from the server before. A watcher will push updates on the server when your history, your playlists or your profiles are updated.
When FreeTube is restarted, history, playlists and profiles will be updated.

```mermaid
sequenceDiagram
    participant Client
    participant Server

    Note over Client, Server: Only once
    Client->>+Server: Send initial history, playlists, profiles
    Server-->>-Client: Response OK/KO

    Note over Client, Server: Before launching FreeTube
    Client->>+Server: Pull update to date history, playlists, profiles
    Server-->>-Client: Response datas
    Client->>+Client: Update databases

    Note over Client, Server: While FreeTube is running
    loop Watch local db updates
        Client->>+Server: Send updated history, playlists, profiles
        Server-->>-Client: Response OK/KO
    end
```

## How to use it?

### Server

To start the server, simply run:

```
/path/to/server
```

### Client

First, sync your local datas to the server:

```
/path/to/client -s http://ip.of.the.server:1323 init
```

Create `~/.bin/freetube-wrapper` and fill it with:

```
#!/bin/sh

/path/to/client -s http://ip.of.the.server:1323 pull
exec /opt/FreeTube/freetube $@
```

Then run `chmod +x ~/.bin/freetube-wrapper`.

Create `~/.local/share/applications/FreeTubeSync.desktop` and fill it with:

```
[Desktop Entry]
Type=Application
Icon=freetube
Terminal=false
Exec=/home/foo/.bin/freetube-wrapper
Name=FreeTube (synced)
```

Create `~/.config/systemd/user/freetubesync-watcher.service` and fill it with:

```
[Unit]
Description=FreeTube Sync Watcher

[Service]
Type=simple
StandardOutput=journal
ExecStart=/path/to/client -s http://ip.of.the.server:1323 watch

[Install]
WantedBy=default.target
```

Then run:

```
systemctl --user daemon-reload
systemctl --user start freetubesync-watcher.service
```

Choose `FreeTube (synced)` to open FreeTube.
