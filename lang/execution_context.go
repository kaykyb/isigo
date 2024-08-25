package lang

import (
	"isigo/context"
	"strings"
)

type ExecutionContext struct {
	context  *context.Context
	children []Node
}

func NewExecutionContext(ctx *context.Context, children []Node) ExecutionContext {
	return ExecutionContext{
		context:  ctx,
		children: children,
	}
}

func (p ExecutionContext) Children() []Node {
	return p.children
}

func (p ExecutionContext) Output() (string, error) {
	var lines []string
	for _, child := range p.children {
		content, err := child.Output()
		if err != nil {
			return "", err
		}

		lines = append(lines, content)
	}

	content := strings.Join(lines, "\n")
	return content, nil
}

func (p ExecutionContext) Eval(ctx *context.Context) (lastReturnVal any, err error) {
	for _, child := range p.children {
		lastReturnVal, err = child.Eval(ctx)
		if err != nil {
			break
		}
	}

	return lastReturnVal, err
}
