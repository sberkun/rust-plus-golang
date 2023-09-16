ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

# PHONY means that it doesn't correspond to a file; it always runs the build commands.

.PHONY: build, build-integrated, run-integrated, clean

build:
	cd rustlib && cargo build --release
	cp rustlib/target/release/libhello.a .
	rm -f main
	go build -o main g1.go main.go

build-integrated:
	cd submission-msm-gpu && cargo build --release
	cp submission-msm-gpu/target/release/libblst_msm.so libblst_msm.so
	rm -f main-integrated
	go build -o main-integrated -ldflags="-r $(ROOT_DIR)" g1.go main-integrated.go


run: build
	@./main

run-integrated: build-integrated
	@./main-integrated

clean:
	cd rustlib && cargo clean
	cd submission-msm-gpu && cargo clean
	rm -rf main main-integrated libhello.a
