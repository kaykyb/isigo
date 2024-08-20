package lang

import (
	"isigo/context"
)

type Block struct {
	context *context.Context
	child   Node
}

func NewBlock(ctx *context.Context, child Node) Block {
	return Block{
		context: ctx,
		child:   child,
	}
}

func (p Block) Output() (string, error) {
	content, err := p.child.Output()
	if err != nil {
		return "", err
	}

	return content, nil
}
