build: build-assembler build-linker build-emulator

build-asm: build-assembler

build-assembler:
	go build -o bin/asme8 assembler/cmd/main.go

build-linker:
	go build -o bin/ld linker/cmd/main.go

build-emulator:
	go build -o bin/emu_asm8 emulator/cmd/main.go

build-tools-and-source: build build-source

build-source:
	@cd source && make && cd ..

clean-source:
	@rm -r source/bin
	@rm -r source/elf

.PHONY: build
