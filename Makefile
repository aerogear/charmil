TAGS="dev"

run:
	go run ./examples/host

build:
	go build -o host_example ./examples/host
	go build -o validator_example ./validator/example

build/tags:
	go build -tags ${TAGS} ./examples/host

test/validator:
	go test ./validator/example