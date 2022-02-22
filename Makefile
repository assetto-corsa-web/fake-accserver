
run-dev:
	go run main.go

build-for-osx:
	go build -o build/osx/accServer.exe main.go

build-for-windows:
	CGO_ENABLED=0 GOOS=windows go build -o build/windows/accServer.exe main.go
