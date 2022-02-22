.PHONY: run build proto
run:
	go run member.go -f etc/member.yaml

build:
	go build .

proto:
	goctl rpc protoc member.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.
	protoc-go-inject-tag -input=./pb/* -XXX_skip=xorm

