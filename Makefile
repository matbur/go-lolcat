IMG_NAME = "lolcat_img"
BIN_NAME = "lolcat"

run: build
	@echo
	@echo "Running app..."
	./$(BIN_NAME) -h | ./$(BIN_NAME)

build:
	@echo
	@echo "Building app..."
	docker run --rm -v $$PWD:/app $(IMG_NAME)
	@echo "App built"

build-image:
	docker build -t $(IMG_NAME) .
