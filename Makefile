# note normally values and secrets like ports and passwords would _not_ be included in VCS and would
# instead be pulled from env vars or a secret manager, hard-coding here for convenience/ease-of-demo purposes

run: export SERVER_ADDRESS=localhost
run: export SERVER_PORT=8080
run: export DB_USER=root
run: export DB_PASSWORD=badpassword123
run: export DB_ADDRESS=localhost
run: export DB_PORT=3306
run: export DB_NAME=portfolio_app

run:
	@go build -o portfolio-app .
	./portfolio-app $(ARGS)

test:
	@go test -v ./...