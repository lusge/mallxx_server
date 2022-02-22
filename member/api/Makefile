.PHONY: run build proto
run:
	go run member.go -f etc/member-api.yaml

build:
	go build .

goctl:
	goctl api go -api member.api -dir . -style gozero

