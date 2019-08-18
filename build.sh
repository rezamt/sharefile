#!/usr/bin/env zsh
rm -f ./sharefile.exe
env GOOS=windows GOARCH=amd64 go build -o target/win64/sharefile.exe