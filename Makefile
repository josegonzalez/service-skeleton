.PHONY: build
build:
	cd examples/hello-world && go mod tidy && go build
