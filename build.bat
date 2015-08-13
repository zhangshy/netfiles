@echo off
go install netfiles\webmain
xcopy %GOPATH%\src\netfiles\template %GOPATH%\bin\template /e
