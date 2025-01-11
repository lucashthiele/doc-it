run:
	@templ generate
	@go mod tidy
	@air

build:
	@templ generate
	@go mod tidy
	@go build -o ./build/program ./cmd