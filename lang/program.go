package lang

import (
	"fmt"
	"isigo/common"
	"isigo/context"
)

type Program struct {
	context *context.Context
	child   Block
}

func NewProgram(ctx *context.Context, child Block) Program {
	return Program{
		context: ctx,
		child:   child,
	}
}

func wrappedProgram(content string) string {
	return fmt.Sprintf(`package main

import (
	"fmt"
)

func scanfPanicInt(to *int) {
	_, err := fmt.Scanf("%%d", to)
	if err != nil {
		panic(err)
	}
}

func scanfPanicFloat(to *float64) {
	_, err := fmt.Scanf("%%f", to)
	if err != nil {
		panic(err)
	}
}


func main() {
%s
}
`, content)
}

func (p Program) Output() (string, error) {
	content, err := p.child.Output()
	if err != nil {
		return "", err
	}

	return wrappedProgram(common.Indent(content)), nil
}
