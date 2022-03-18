gopath = $(shell go env GOPATH)
basedir = $(gopath)/src/mallxx_server

.PHONY: run build proto
run:
	go run order.go -f etc/order.yaml

build:
	go build .

proto:
	goctl rpc protoc $(basedir)/order/rpc/order.proto --proto_path=$(basedir) --go_out=./ --go-grpc_out=./ --zrpc_out=.
	protoc-go-inject-tag -input=./pb/* -XXX_skip=xorm

