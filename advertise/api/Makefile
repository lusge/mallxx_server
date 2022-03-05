.PHONY: run build proto
run:
	go run advertise.go -f etc/advertise-api.yaml

build:
	go build .

goctl:
	goctl api go -api advertise.api -dir . -style gozero

