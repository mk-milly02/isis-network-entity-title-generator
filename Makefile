.PHONY: build

build:
	go build -o ipv4tonet.exe .\cmd\main.go
test:
	go test 
run:
	go run cmd/*.go