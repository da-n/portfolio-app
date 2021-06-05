run:
	@go build -o portfolio-app .
	./portfolio-app $(ARGS)

test:
	@go test ./...