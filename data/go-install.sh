#!/bin/bash
apt-get update
apt-get install -y git
wget 'https://go.dev/dl/go1.18.linux-amd64.tar.gz'
tar -C '/usr/local' -xzf 'go1.18.linux-amd64.tar.gz'
export PATH="$PATH:/usr/local/go/bin"
export GOPATH="/app"
cd "${GOPATH}/src/tjweldon/pdf-printer" || exit 1
go install
