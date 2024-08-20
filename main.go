package main

import (
	"flag"
	"fmt"
	"isigo/context"
	"isigo/lexer"
	"isigo/parser"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("./input_dec.isi")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	l := lexer.LexicalAnalysis{}
	l.SetContent(string(content))

	p := parser.New(&l)
	ctx := context.New()
	prog, _, err := p.Prog(&ctx)

	fmt.Println(prog)
	fmt.Println(err)

	main_alt()
}

func main_alt() {
	// Define o argumento de linha de comando
	filePath := flag.String("file", "", "O caminho para o arquivo de entrada")
	flag.Parse()

	if *filePath == "" {
		fmt.Println("Uso: isigo -file=<caminho_para_arquivo>")
		os.Exit(1)
	}

	// Lê o arquivo
	content, err := os.ReadFile(*filePath)
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	l := lexer.LexicalAnalysis{}
	l.SetContent(string(content))

	p := parser.New(&l)
	ctx := context.New()
	prog, _, err := p.Prog(&ctx)

	// Imprime o resultado
	fmt.Println(prog)
	fmt.Println(err)
}
