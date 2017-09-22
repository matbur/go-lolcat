BIN_NAME = "lolcat"

run: build
	@cat mono | ./$(BIN_NAME)

build:
	@docker run --rm -v $$PWD:/app -w /app golang:alpine \
		go build -o $(BIN_NAME)
