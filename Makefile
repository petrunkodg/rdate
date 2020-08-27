GO=go
GOPRIVATE?=gitlab.com/*

modules:
	GOPRIVATE=$(GOPRIVATE) $(GO) mod tidy -v
	$(GO) mod verify

test: modules
	$(GO) test -race -v -cover ./...

coverage: modules
	$(GO) test -race -v -covermode=atomic -coverpkg=./... -coverprofile=coverage.out ./...