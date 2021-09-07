#!/bin/bash

echo 'Clone and build...'
git clone https://github.com/overstimate/SniperWUI
git checkout stable
go build server.go
echo 'Install server to /usr/local/bin'
echo '> sudo install server /usr/local/bin'
sudo install server /usr/local/bin
echo 'Done. Start the server using the command below when ready:'
echo 'PORT=8080 server'