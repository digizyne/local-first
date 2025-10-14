default: fmt run

fmt:
    go fmt ./...

run:
    go run main.go

build:
    go build -o lf main.go