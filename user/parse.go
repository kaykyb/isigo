package user

import (
	"isigo/context"
	"isigo/lang"
	"isigo/lexer"
	"isigo/parser"
	"log"
	"os"
)

func ParseFromFile(filePath string) lang.Program {
	// Lê o arquivo
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	l := lexer.LexicalAnalysis{}
	err = l.SetContent(string(content))

	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	p := parser.New(&l)
	ctx := context.New()
	prog, delta, err := p.Prog(&ctx)
	if err != nil {
		log.Fatalf("Erro: [Linha %d, Coluna %d]: %v", delta.Position().Line+1, delta.Position().Column+1, err)
	}

	return prog
}
