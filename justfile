default: install

fmt:
    go fmt ./...

run:
    go run main.go

build:
    go build -o lf main.go

tidy:
    go mod tidy

add PACKAGE:
    go get -u {{PACKAGE}}

install:
    go install