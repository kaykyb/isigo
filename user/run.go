package user

import (
	"fmt"
	"isigo/context"
	"isigo/lang"
)

func Run(prog lang.Program) {
	newContext := context.New()
	val, err := prog.Eval(&newContext)

	if err != nil {
		panic(err)
	}

	fmt.Println("-> ", val)
	fmt.Println("[ Programa encerrado. ]")
}
