.PHONY: run build proto
run:
	go run cart.go -f etc/cart-api.yaml

build:
	go build .

goctl:
	goctl api go -api cart.api -dir . -style gozero

