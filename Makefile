dev:
	go run cmd/cli/main.go
build:
	go build -o bin/cynotes cmd/cli/main.go
install:
	go install cmd/cli/main.go


