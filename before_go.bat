@echo off

cd .\sl
go mod init sl
go mod tidy

cd ..\cmdutil
go mod init cmdutil
go mod edit -replace sl=../sl
go mod tidy

cd ..

set "batchFilePath=%~dp0"
for %%F in ("%batchFilePath%.") do (
    set "folderName=%%~nxF"
)

go mod init %folderName%
go mod edit -replace sl=./sl
go mod edit -replace cmdutil=./cmdutil
go mod tidy