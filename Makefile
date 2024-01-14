
test: 
	go test

lint:
	docker run --rm -v $(shell pwd):/workspace -w /workspace golangci/golangci-lint:latest golangci-lint run