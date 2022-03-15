.PHONY: run build proto
run:
	go run cart.go -f etc/product.yaml

build:
	go build .

proto:
	goctl rpc protoc cart.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.
	protoc-go-inject-tag -input=./pb/* -XXX_skip=xorm

