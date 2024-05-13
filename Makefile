.PHONY: build lint test generate

lint:
	golangci-lint run

test:
	go test -v ./...

generate:
	protoc -I api api/netvuln/netvuln.proto --go_out=./pkg --go-grpc_out=./pkg --go-grpc_opt=require_unimplemented_servers=false  