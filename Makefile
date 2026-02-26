build: build-assembler build-linker build-emulator

test:
	go clean -testcache
	go test -v ./assembler/test/...
	go test -v ./emulator/test/...

build-asm: build-assembler

build-assembler:
	go build -o bin/asme8 assembler/cmd/main.go

build-linker:
	go build -o bin/ld linker/cmd/main.go

build-emulator:
	go build -o bin/emu_asme8 emulator/cmd/main.go

.PHONY: source
source: build build-source

build-source: clean-source
	@cd source && make

clean-source:
	@rm -rf source/bin
	@rm -rf source/elf

.PHONY: build

run:
	go run emulator/cmd/main.go --config source/linker_config --load-bin source/bin/source.bin --symbol-file source/bin/source.sym --delay 1ns
