.DEFAULT_GOAL := build
HOST_BIN_FILE = host

build:
	@go build -o "${HOST_BIN_FILE}" ./examples/"${HOST_BIN_FILE}"
run_host:
	./"${HOST_BIN_FILE}"
clean:
	go clean
	rm -f ${HOST_BIN_FILE}