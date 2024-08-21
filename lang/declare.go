package lang

import (
	"fmt"
	"isigo/context"
	"strings"
)

type Declare struct {
	context   *context.Context
	variables []Variable
}

func NewDeclare(ctx *context.Context, variables []Variable) Declare {
	return Declare{
		context:   ctx,
		variables: variables,
	}
}

func (p Declare) Output() (string, error) {
	var lines []string
	for _, variable := range p.variables {
		lines = append(lines, fmt.Sprintf("var %s %s", variable.identifier, variable.variableType.Output()))
	}

	content := strings.Join(lines, "\n")
	return content, nil
}
