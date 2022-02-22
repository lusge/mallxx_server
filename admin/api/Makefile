.PHONY: run build proto
run:
	go run admin.go -f etc/admin-api.yaml

build:
	go build .

goctl:
	goctl api go -api admin.api -dir . -style gozero

