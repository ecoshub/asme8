build: object-files link

object-files:
	@mkdir -p elf
	@cd elf && find ../lib -type f | grep ".asm" | xargs -I {} sh -c '../../bin/./asme8 --mode elf --file {}'

link: 
	@mkdir -p bin
	@cd bin && ../../bin/./ld --memory-config ../memory_config --segment-config ../segment_config --print --output link.bin ../elf/*.o

.PHONY: build
