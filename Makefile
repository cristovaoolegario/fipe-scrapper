lint: 
		go mod tidy
		go fmt ./... 
		go vet ./...

test:
	go test ./... -v