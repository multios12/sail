echo off

rem build front page
call yarn build
robocopy .\build .\srv\static /MIR
rmdir /s /Q build

rem build server app
set GOOS=linux
set GOARCH=amd64
mkdir dist
cd convert
go build -o ../dist
cd ..
cd srv
go build -o ../dist 
cd ..