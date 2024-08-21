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
	"bufio"
	"os"
	"strconv"
)

func scanLine() string {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return line
}

func scanfPanicInt() int64 {
	fval, err := strconv.ParseInt(scanLine(), 10, 64)
	
	if err != nil { 
		panic(err)
	}

	return fval
}

func scanfPanicFloat() float64 {
	fval, err := strconv.ParseFloat(scanLine(), 64)
	
	if err != nil { 
		panic(err)
	}

	return fval
}

func scanfPanicString() string {
	return scanLine()
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
