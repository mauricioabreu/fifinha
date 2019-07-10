build:
	go build -o fifinha

run: build
	./fifinha

.PHONY: build run