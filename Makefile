GOOS ?= darwin
GOARCH ?= amd64
VERSION ?= latest

# 如果是运行在 mac m1 芯片下的docker中，则需要将 goarch改为 arm64 而不是默认的 amd64
build:
	# @touch *
	@echo "Buildding cms VERSION=$(VERSION)"
	@go run *.go
	@mkdir -p ./compiled
	@cp -r ./config.json ./compiled/config.json
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X main.version=$(VERSION)" -o ./compiled/imgArrange.exe *.go
	@echo "Success!"