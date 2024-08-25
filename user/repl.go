package user

import (
	"bufio"
	"fmt"
	"isigo/context"
	"isigo/lexer"
	"isigo/parser"
	"isigo/sources"
	"os"
	"strings"
)

func Repl() {
	fmt.Println("ISIGO REPL")

	initialCtx := context.New()
	currentCtx := &initialCtx

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("isigo> ")
		nextCommand, _ := reader.ReadString('\n')
		nextCommandCompiled := fmt.Sprintf("programa\n%sfimprog.", nextCommand)
		source := sources.NewBuildReader(bufio.NewReader(strings.NewReader(nextCommandCompiled)))

		l := lexer.New(source)
		p := parser.NewReplParser(&l)

		prog, _, err := p.ParseProgram(currentCtx)

		if err != nil {
			fmt.Printf("🔴 %v\n", err)
			continue
		}

		result, err := prog.Eval(currentCtx)
		if err != nil {
			fmt.Printf("🔴 %v\n", err)
			continue
		}

		fmt.Printf("🟢 %v\n", result)

		currentCtx = prog.DeepestContext()
	}
}
