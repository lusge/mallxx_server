.PHONY: run build proto
run:
	go run advertise.go -f etc/advertise.yaml

build:
	go build .

proto:
	goctl rpc protoc advertise.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.
	protoc-go-inject-tag -input=./pb/* -XXX_skip=xorm

