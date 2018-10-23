all: compile zip

compile:
	GOOS=linux GOARCH=amd64 go build -o main main.go

zip:
	zip main.zip main

