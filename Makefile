dev:
	@set -e; \
	wire; \
	air -c .air.toml

build:
	wire
	go build -o bin/server .

run: build
	 ./bin/server