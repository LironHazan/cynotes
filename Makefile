dev:
	go run cmd/cli/main.go
build:
	go build -o bin/cynotes cmd/cli/main.go
install:
	go install cmd/cli/main.go # should fix this, the bin won't be copy to $GOPATH/bin --> https://stackoverflow.com/questions/24069664/what-does-go-install-do


