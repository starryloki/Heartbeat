mkdir -p bin
n="Heartbeat"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/${n}_Linux_amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/${n}_Darwin_amd64
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/${n}_Darwin_arm64
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o bin/${n}_FreeBSD_amd64
CGO_ENABLED=0 GOOS=netbsd GOARCH=amd64 go build -o bin/${n}_NetBSD_amd64
CGO_ENABLED=0 GOOS=openbsd GOARCH=amd64 go build -o bin/${n}_OpenBSD_amd64
CGO_ENABLED=0 GOOS=solaris GOARCH=amd64 go build -o bin/${n}_Solaris_amd64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/${n}_Windows_amd64.exe
