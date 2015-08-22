@echo off
::go install netfiles\webmain
::xcopy %GOPATH%\src\netfiles\template_webmain %GOPATH%\bin\template /e
:: "||"作用，当||前命令执行失败时才执行||后的命令
go install netfiles\pcshow || pause
xcopy %GOPATH%\src\netfiles\static %GOPATH%\bin\static /e /Y
