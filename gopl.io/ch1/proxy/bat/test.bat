REM @echo on
REM title TestBAT
REM echo This is the First Class
REM echo.|echo printf
REM echo =========

REM :flag
REM set /a var+=1
REM echo %var%
REM if %var% leq 43 goto flag
REM echo.
start C:\winAPIServer.exe
start cmd /k "cd/d C:\ && winWebServer.exe"
