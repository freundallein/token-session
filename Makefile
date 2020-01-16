init:
	git config core.hooksPath .githooks
run:
	go run docs/examples/example.go
test:
	go test -cover ./...
