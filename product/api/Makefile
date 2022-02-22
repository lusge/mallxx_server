.PHONY: run build proto
run:
	go run product.go -f etc/product-api.yaml

build:
	go build .

goctl:
	goctl api go -api product.api -dir . -style gozero

