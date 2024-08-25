package user

import (
	"bufio"
	"isigo/context"
	"isigo/lang"
	"isigo/lexer"
	"isigo/parser"
	"isigo/sources"
	"log"
	"os"
)

func ParseFromFile(filePath string) lang.Program {
	// Lê o arquivo
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	sourceReader := sources.NewBuildReader(reader)

	l := lexer.New(sourceReader)
	p := parser.New(&l)
	ctx := context.New()
	prog, delta, err := p.ParseProgram(&ctx)
	if err != nil {
		log.Fatalf("Erro: [Linha %d, Coluna %d]: %v", delta.Position().Line+1, delta.Position().Column+1, err)
	}

	return prog
}
