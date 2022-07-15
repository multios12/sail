#!/bin/sh

export HTML="sail.html"
yarn --cwd ./front build
cp ./front/dist/* ./server/static/ -R
export HTML="memo.html"
yarn --cwd ./front build
cp ./front/dist/* ./server/static/ -R
export HTML="diary.html"
yarn --cwd ./front build
cp ./front/dist/* ./server/static/ -R
find ./front/dist -delete

cd server
GOOS=linux
GOARCH=amd64
go build -ldflags="-s -w" -trimpath -o ../dist/
cd ..