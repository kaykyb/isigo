package main

import (
	"flag"
	"fmt"
	"isigo/user"
	"os"
)

func main() {
	command := flag.String("c", "", "Comando a ser executado: build, run, repl")
	filePath := flag.String("f", "", "O caminho para o arquivo de entrada")
	flag.Parse()

	if *command == "" || (*command != "build" && *command != "run" && *command != "repl") {
		fmt.Println("Uso: isigo -c=<build|run|repl> -f=<caminho_para_arquivo>")
		os.Exit(1)
	}

	if *command != "repl" && *filePath == "" {
		fmt.Println("Uso: isigo -c=<build|run> -f=<caminho_para_arquivo>")
		os.Exit(1)
	}

	switch *command {
	case "build":
		build(*filePath)
	case "run":
		run(*filePath)
	default:
		repl()
	}
}

func build(filePath string) {
	prog := user.ParseFromFile(filePath)
	user.Build(prog)
}

func run(filePath string) {
	prog := user.ParseFromFile(filePath)
	user.Run(prog)
}

func repl() {
	user.Repl()
}
