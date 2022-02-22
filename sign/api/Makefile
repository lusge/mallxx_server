.PHONY: run build proto
run:
	go run sign.go -f etc/sign-api.yaml

build:
	go build .

goctl:
	goctl api go -api sign.api -dir . -style gozero

