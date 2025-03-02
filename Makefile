run:
	@go mod tidy
	@air

build:
	@go mod tidy
	@go build -o ./build/program ./cmd/api