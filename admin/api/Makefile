.PHONY: run build proto
run:
	go run admin.go -f etc/admin-api.yaml

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o admin .

goctl:
	goctl api go -api admin.api -dir . -style gozero

