@echo off

rem 真实路径
set RootFolderPath=E:\ncManage
set SourceFolder=%RootFolderPath%\src\Presentation\NCManageUI
set DeployFolder=D:\deployFolder
rem set DestExePath=%SourceFolder%\bin

set MsBuildExePath=C:\Windows\Microsoft.NET\Framework\v4.0.30319
set ServiceName=SMS_Service

set InstallUtilPath=C:\Windows\Microsoft.NET\Framework\v4.0.30319


rem pull最新代码
REM cd %RootFolderPath%
REM git pull

rem 进入项目
cd %SourceFolder%

rem 编译项目
REM net stop %ServiceName%
REM C:\soft\nuget.exe restore %SourceFolder%\SMS_sln.sln
REM %MsBuildExePath%\MSBuild.exe %SourceFolder%\SMS_sln.sln /p:Configuration=Release /p:TargetFramework=v4.5.2 /p:Platform="Any CPU"
start npm run build

rem 部署项目
REM echo --------------
REM xcopy %SourceFolder%\bin\*  %DeployFolder%  /y /e /I /r
mstsc /v: 10.10.0.123 /console
net use \\10.10.0.123\ipc$ Nc@test /user:Administrator
xcopy %SourceFolder%\dist\static \\10.10.0.123\D\test /y
xcopy %SourceFolder%\dist\index.html \\10.10.0.123\D\test /y
net use \\10.10.0.123\ipc$ /delete

rem 停止，卸载服务
REM net stop %ServiceName%
REM %InstallUtilPath%\InstallUtil.exe -u %DeployFolder%\Release\SMS_Service.exe

rem 安装，启动服务
echo %InstallUtilPath%\InstallUtil.exe %DeployFolder%\Release\SMS_Service.exe
%InstallUtilPath%\InstallUtil.exe %DeployFolder%\Release\SMS_Service.exe

echo net start %ServiceName%
net start %ServiceName%

pause