#!/bin/sh

export HTML="index.html"
yarn --cwd ./front build
cp ./front/dist/* ./server/static/ -R

cd server
GOOS=linux
GOARCH=amd64
go build -ldflags="-s -w" -trimpath -o ../dist/
cd ..