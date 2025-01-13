out:="parser"

# make generate path=AsmE8.g4 out=src/parser
generate:
	/opt/homebrew/bin/antlr4 -Dlanguage=Go $(path) -o $(out)

run:
	go run cmd/main.go $(file)
