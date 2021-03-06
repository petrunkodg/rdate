GO=go

modules:
	$(GO) mod tidy -v
	$(GO) mod verify

test: modules
	$(GO) test -race -v -cover ./...

coverage: modules
	$(GO) test -race -v -covermode=atomic -coverpkg=./... -coverprofile=coverage.out ./...
	cat coverage.out >> coverage.txt