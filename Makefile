APP_NAME=main
BINARY=./$(APP_NAME)
SRC=main.go

.PHONY: all build run clean

all: build run clean

build:
	@go build -o $(APP_NAME) $(SRC)

run:
	@$(MAKE) build
	@$(BINARY) $(path) $(col)

clean:
	@rm -f $(APP_NAME)
