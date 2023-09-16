ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

# PHONY means that it doesn't correspond to a file; it always runs the build commands.

.PHONY: build
build:
	cd rustlib && cargo build --release
	cp rustlib/target/release/libhello.a .
	rm -f main
	go build -o main g1.go main.go

.PHONY: run
run: build
	@./main

.PHONY: clean
clean:
	cd rustlib && cargo clean
	rm -rf main libhello.a
