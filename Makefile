# note normally values and secrets like ports and passwords would _not_ be included in VCS and would
# instead be pulled from env vars or a secret manager, hard-coding here for convenience/ease-of-demo purposes

run: export SERVER_ADDRESS=localhost
run: export SERVER_PORT=8080
run:
	@go build -o portfolio-app .
	./portfolio-app $(ARGS)

test:
	@go test -v ./...