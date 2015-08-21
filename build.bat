@echo off
::go install netfiles\webmain
::xcopy %GOPATH%\src\netfiles\template_webmain %GOPATH%\bin\template /e
go install netfiles\pcshow
xcopy %GOPATH%\src\netfiles\static %GOPATH%\bin\static /e
