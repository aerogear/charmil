TAGS="dev"

run:
	go run ./examples/host

build:
	go build ./examples/host

build/tags:
	go build -tags ${TAGS} ./examples/host

test:
	go test ./validator/example