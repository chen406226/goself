set dirs="C:\Users\Administrator\AppData\Roaming\JetBrains"
cd dirs
REM rmdir "eval" /s /q
REM del "options\other.xml"
@Echo off
for /f "delims=" %%a in ('dir /ad /b /s C:\Users\Administrator\AppData\Roaming\JetBrains') do (
echo %%a|find "eval" >nul && (rmdir %%a /s /q & echo)
)
@Echo off
REM for /f "delims=" %%i in ('dir /ad/b/s "%dirs%"') do (echo %%i|find "options" >nul && echo %%i)
REM for /f "delims=" %%i in ('dir /a-d/b/s "%dirs%*other.xml*"') do (echo %%i)
for /f "delims=" %%i in ('dir other.xml /b /s') do del %%i
reg delete "HKEY_CURRENT_USER\Software\JavaSoft\Prefs\jetbrains\idea" /f
reg delete "HKEY_CURRENT_USER\Software\JavaSoft\Prefs\jetbrains\rider" /f
reg delete "HKEY_CURRENT_USER\Software\JavaSoft\Prefs\jetbrains\webstorm" /f
