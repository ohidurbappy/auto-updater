mkdir -p bin
go build -ldflags="-s -w" -o bin/auto-updater.exe main.go
tar -cvzf bin/auto-updater.tar.gz bin/auto-updater.exe