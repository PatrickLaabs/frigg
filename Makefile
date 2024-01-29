.PHONY: build-darwin-amd64
build-darwin-amd64:
	go build -o argohub main.go


#build-darwin-arm:
#    GOOS=darwin
#    GOARCH=arm
#	go build -o $(BINARY_NAME)-${GOOS}-${GOARCH} main.go
#
#build-linux:
#    GOOS=linux
#    GOARCH=amd64
#	go build -o $(BINARY_NAME)-${GOOS}-${GOARCH} main.go
#
#build-windows:
#    GOOS=windows
#    GOARCH=amd64
#	go build -o $(BINARY_NAME)-${GOOS}-${GOARCH} main.go
#
#clean:
#	rm -f $(BINARY_NAME)-darwin-amd64 $(BINARY_NAME)-darwin-arm $(BINARY_NAME)-linux-amd64 $(BINARY_NAME)-windows-amd64.exe
