out:="src/parser"
path:="AsmE8.g4"
antlr4_path:="/opt/homebrew/bin/antlr4"

generate:
	$(antlr4_path) -Dlanguage=Go $(path) -o $(out)

run:
	go run cmd/main.go $(file)

.PHONY: test
test:
	go test -v ./test/unit_test/...
