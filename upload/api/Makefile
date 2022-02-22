.PHONY: run build proto
run:
	go run upload.go -f etc/upload-api.yaml

build:
	go build .

goctl:
	goctl api go -api upload.api -dir . -style gozero

