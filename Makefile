.DEFAULT_GOAL := build
CHARMIL_BIN_FILE = charmil
HOST_BIN_FILE = hostCLI

build:
	@go build -o "${CHARMIL_BIN_FILE}" ./cmd/"${CHARMIL_BIN_FILE}"
	@go build -o "${HOST_BIN_FILE}" ./examples/"${HOST_BIN_FILE}"
run_charmil:
	./"${CHARMIL_BIN_FILE}"
run_host:
	./"${HOST_BIN_FILE}"
clean:
	go clean
	rm -f ${CHARMIL_BIN_FILE} ${HOST_BIN_FILE}