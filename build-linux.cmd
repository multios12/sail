echo off

rem build front page
cd front
call yarn build
robocopy .\build .\server\static /MIR
rmdir /s /Q build
cd ..

rem build server app
set GOOS=linux
set GOARCH=amd64
cd server
go build -ldflags="-s -w" -trimpath -o ../dist/
cd ..