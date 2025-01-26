build: object-files link

object-files:
	@echo 'elf assemble started'
	@mkdir -p elf
	@cd elf && find ../lib -type f | xargs -I {} sh -c '../../bin/./asme8 --mode elf --file {}'
	@echo 'elf assemble done'

link: 
	@echo 'linking started'
	@mkdir -p bin
	@cd bin && ../../bin/./ld --config ../linker_config --output link.bin ../elf/*.o
	@echo 'linking done'

.PHONY: build
