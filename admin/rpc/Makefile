.PHONY: run build proto
run:
	go run admin.go -f etc/admin.yaml

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o admin .

proto:
	goctl rpc protoc admin.proto --go_out=./ --go-grpc_out=./ --zrpc_out=.
	protoc-go-inject-tag -input=./pb/* -XXX_skip=xorm

