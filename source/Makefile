build: object-files link

object-files:
	@mkdir -p elf
	@cd elf && find ../lib -type f | grep ".asm" | xargs -I {} sh -c '../../bin/./asme8 --mode elf --file {}'

link: 
	@mkdir -p bin
	@cd bin && ../../bin/./ld --config ../linker_config --output link.bin ../elf/*.o

.PHONY: build
