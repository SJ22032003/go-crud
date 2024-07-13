run:
	@air -c .air.toml

install:
	@echo "Installing dependencies"
	@go mod download
	@go mod tidy
	@go mod vendor 
