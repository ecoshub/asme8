build: build-assembler build-linker build-emulator

test:
	go test ./emulator/test/...
	go test ./assembler/test/...

build-asm: build-assembler

build-assembler:
	go build -o bin/asme8 assembler/cmd/main.go

build-linker:
	go build -o bin/ld linker/cmd/main.go

build-emulator:
	go build -o bin/emu_asme8 emulator/cmd/main.go

build-tools-and-source: build build-source

build-source:
	@cd source && make && cd ..

clean-source:
	@rm -rf source/bin
	@rm -rf source/elf

.PHONY: build
