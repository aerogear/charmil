.DEFAULT_GOAL := build
BIN_FILE = charmil

build:
	@go build -o "${BIN_FILE}" ./cmd/"${BIN_FILE}"
run:
	./"${BIN_FILE}"
clean:
	go clean
	rm -f ${BIN_FILE}