BIN_DIR=_output/bin

all: ssh-board add

init:
	mkdir -p ${BIN_DIR}

ssh-board: init
	go build -v -o=${BIN_DIR}/ssh-board ./cmd/ssh-board/

add: init
	go build -v -o=${BIN_DIR}/add ./cmd/add/

clean:
	rm -rf _output/