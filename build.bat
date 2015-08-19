@echo off
::go install netfiles\webmain
::xcopy %GOPATH%\src\netfiles\template_webmain %GOPATH%\bin\template /e
go install netfiles\pcshow
xcopy %GOPATH%\src\netfiles\template %GOPATH%\bin\template /e
