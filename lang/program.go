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

func (p Program) Child() Block {
	return p.child
}

func wrappedProgram(content string) string {
	return fmt.Sprintf("package main\n\nimport \"isigoprogram/std\"\n\nfunc main() {\n%s\n}\n", content)
}

func (p Program) Output() (string, error) {
	content, err := p.child.Output()
	if err != nil {
		return "", err
	}

	return wrappedProgram(common.Indent(content)), nil
}

func (p Program) Eval(ctx *context.Context) (any, error) {
	return p.child.Eval(ctx)
}

func (p Program) DeepestContext() *context.Context {
	if p.child.DeepestContext() != nil {
		return p.child.DeepestContext()
	}

	return p.context
}
