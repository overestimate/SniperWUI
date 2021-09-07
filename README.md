# SniperWUI
Self-hosted highly configurable block game sniper web UI for all.

# Why?
Fuck cowboy.rocks and sniping.fun, we want SELF-HOSTED panels.

# Web UI Usage
See </WEB.md>

# Server functionality
Uses `net/http` to allow for web hosting, with handlers being loaded from `handlers.go`.

# API Documentation
See </API.md>.

# Server Setup
Requires Go 1.16 or newer and git. Instructions require screen, 
although it is not required to run the program.

## Quick Setup
These instructions require cURL to be installed. You may view the
script which is being ran [here](https://github.com/overstimate/SniperWUI/raw/stable/install.sh).

Log in as the user you will be running the server on.

(Optional) Make a new screen with `screen -S swui`.

Run `curl -L https://github.com/overstimate/SniperWUI/raw/stable/install.sh | bash` and type in your password when prompted to elevate to root.

Run the following to start the server:
```
PORT=8080 server
```

If you used `screen`, detach using ^AD and reattach using `screen -r swui`.

## Manual Setup

Run the following commands in a bash shell as root until stated in the comment.
```bash
git clone https://github.com/overstimate/SniperWUI
git checkout stable
go build server.go
install server /usr/local/bin
# log in to the user which will be hosting the server
screen -S swui
PORT=8080 server
```

Check screen with `screen -r swui`.