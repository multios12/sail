echo off

rem build front page
call yarn build
robocopy .\build .\srv\static /MIR
rmdir /s /Q build

rem build server app
set GOOS=linux
set GOARCH=amd64
mkdir dist
cd srv
go build -ldflags="-s -w" -trimpath -o ../dist 
cd ..