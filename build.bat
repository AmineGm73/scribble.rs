@echo off
del .\scribblers.exe 2>nul
go build ./cmd/scribblers
.\scribblers