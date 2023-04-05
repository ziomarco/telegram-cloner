.DEFAULT_GOAL := build

build:
	make clean
	@echo "Building..."
	go build
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o dist/tcl-amd64.exe main.go
	CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -o dist/tcl-386.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o dist/tcl-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o dist/tcl-darwin-arm64 main.go
	GOOS=linux GOARCH=386 go build -o dist/tcl-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o dist/tcl-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o dist/tcl-linux-arm64 main.go
	rm -rf mobile-security-hashgenerator
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