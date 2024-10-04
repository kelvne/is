GOBIN ?= $$(go env GOPATH)/bin

.PHONY: prepare install-tools test test-coverage check-coverage

prepare: install-tools

install-tools:
	@go install github.com/vladopajic/go-test-coverage/v2@latest

test:
	@go test

test-coverage:
	@go test ./... -coverprofile=./cover.out -covermode=atomic -coverpkg=./...

check-coverage: test-coverage
	@${GOBIN}/go-test-coverage --config=./.testcoverage.yml