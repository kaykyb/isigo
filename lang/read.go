package lang

import (
	"bufio"
	"fmt"
	"isigo/context"
	"isigo/symbol"
	"isigo/value_types"
	"os"
	"strconv"
)

type Read struct {
	context *context.Context
	output  *symbol.Symbol
}

func NewRead(ctx *context.Context, output *symbol.Symbol) Read {
	return Read{
		context: ctx,
		output:  output,
	}
}

func (p Read) Output() (string, error) {
	switch p.output.Type {
	case value_types.IntegerValueTypeEntity:
		return fmt.Sprintf("%s = scanfPanicInt()", p.output.Identifier), nil
	case value_types.FloatValueTypeEntity:
		return fmt.Sprintf("%s = scanfPanicFloat()", p.output.Identifier), nil
	default:
		return fmt.Sprintf("%s = scanfPanicString()", p.output.Identifier), nil
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

func (p Read) Eval(ctx *context.Context) (any, error) {
	var val any

	switch p.output.Type {
	case value_types.IntegerValueTypeEntity:
		val = scanfPanicInt()
	case value_types.FloatValueTypeEntity:
		val = scanfPanicFloat()
	default:
		val = scanfPanicString()
	}

	p.output.AssignRuntimeValue(val)
	return val, nil
}
