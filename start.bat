@echo off

go build -ldflags "-s -w" -o myprogram.exe ./src
pause