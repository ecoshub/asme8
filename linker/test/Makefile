build: object-files link

object-files:
	@mkdir -p elf
	@cd elf && find .. -type f | grep ".asm" | xargs -I {} sh -c 'go run ../../../assembler/cmd/main.go --mode elf --file {}'

link: 
	@mkdir -p bin
	@cd bin && go run ../../../linker/cmd/main.go --config ../config --print --output source.bin ../elf/*.o

run:
	@go run ../../emulator/cmd/main.go --config config --load-bin bin/source.bin --symbol-file bin/source.sym --delay 1ms

.PHONY: build
