.DEFAULT_GOAL := build

build:
	make clean
	@echo "Building..."
	go build
	GOOS=darwin GOARCH=amd64 go build -o dist/tcl-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o dist/tcl-darwin-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o dist/tcl-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o dist/tcl-linux-arm64 main.go
	rm -rf telegramcloner
clean:
	@echo "Cleaning up"
	rm -rfv dist
core:
	@echo "Starting Telegram Core"
	docker-compose up -d
core-logs:
	docker-compose logs -f --tail 100
run:
	dist/tcl-darwin-arm64 $(ARGS)