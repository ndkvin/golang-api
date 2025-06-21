dev:
	wire
	air -c .air.toml

build:
	wire
	go build -o bin/server .

run:
	 ./bin/server