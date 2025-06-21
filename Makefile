

dev:
	wire
	air -c .air.toml

build:
	go build -o bin/server main.go