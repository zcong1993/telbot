generate:
	@go generate ./...

build: generate
	@echo "====> Build telbot"
	@sh -c ./build.sh
