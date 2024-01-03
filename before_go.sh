#!/bin/bash

echo "running: sudo apt install libncurses5-dev"
sudo apt install libncurses5-dev

cd ./sl
go mod init sl
go mod tidy

cd ../cmdutil
go mod init cmdutil
go mod edit -replace sl=../sl
go mod tidy

cd ..
go mod init ${PWD##*/}
go mod edit -replace sl=./sl
go mod edit -replace cmdutil=./cmdutil
go mod tidy

file_path=$(go env GOMODCACHE)/seehuhn.de/go/ncurses@v0.2.0/keys.go
sudo sed -i "171s/.*/  	\/\/ C.KEY_EVENT:     KeyEvent,/" "$file_path"