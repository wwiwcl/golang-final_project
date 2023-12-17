#!/bin/bash

cd ./sl
go mod init sl
go mod tidy

cd ../cmdutil
go mod init cmdutil
go mod edit -replace sl=../sl
go mod tidy

cd ../
go mod init %PWD
go mod edit -replace cmdutil=./cmdutil
go mod tidy