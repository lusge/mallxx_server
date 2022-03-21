.PHONY: run build proto
run:
	go run sign.go -f etc/order-api.yaml

build:
	go build .

goctl:
	goctl api go -api order.api -dir . -style gozero

