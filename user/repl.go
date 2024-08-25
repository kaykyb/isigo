package user

import (
	"bufio"
	"fmt"
	"isigo/context"
	"isigo/lexer"
	"isigo/parser"
	"os"
)

func Repl() {
	fmt.Println("--- ISIGO REPL ----")

	ctx := context.New()

	for {
		command := readCommand()

		l := lexer.New("programa " + command + " fimprog.")
		p := parser.New(&l)
		prog, _, err := p.ParseProgram(&ctx)

		if err != nil {
			fmt.Printf("🔴 %v\n", err)
			continue
		}

		result, err := prog.Eval(&ctx)
		if err != nil {
			fmt.Printf("🔴 %v\n", err)
			continue
		}

		fmt.Printf("🟢 %v\n", result)
	}
}

func scanLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return line[:len(line)-1]
}

func readCommand() string {
	fmt.Print("👉 ")
	return scanLine()
}
