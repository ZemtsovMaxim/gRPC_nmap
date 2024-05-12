.PHONY: build lint test

build:
	go build -o netvuln-service main.go

lint:
	golangci-lint run

test:
	go test -v ./...

generate:
	protoc -I api api/netvuln/v1/netvuln.proto --go_out=./pkg --go-grpc_out=./pkg --go-grpc_opt=require_unimplemented_servers=false  