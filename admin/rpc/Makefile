.PHONY: run build proto
run:
	go run admin.go -f etc/admin.yaml

build:
	go build .

proto:
	goctl rpc protoc admin.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.
	protoc-go-inject-tag -input=./pb/* -XXX_skip=xorm

