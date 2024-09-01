package parser

import (
	"isigo/context"
	"isigo/failure"
	"isigo/lang"
	"isigo/syntax"
	"isigo/tokens"
	"isigo/value_types"
	"strconv"
)

func (c *Parser) Expr(ctx *context.Context, delta TokenDelta) (lang.Expr, TokenDelta, error) {
	factor, delta, err := c.Factor(ctx, delta)
	if err != nil {
		return lang.TermExpr{}, delta, err
	}

	return c.exprInternal(ctx, factor, delta)
}

func (c *Parser) exprInternal(ctx *context.Context, left lang.Factor, delta TokenDelta) (lang.Expr, TokenDelta, error) {
	if delta.token.IsOperator() && (delta.token.Is(syntax.Sum) || delta.token.Is(syntax.Minus)) {
		factorTerm := lang.NewFactorTerm(ctx, left)
		termExpr := lang.NewTermExpr(ctx, factorTerm)

		return c.exprAux(ctx, termExpr, delta)
	}

	term, delta, err := c.term(ctx, left, delta)
	if err != nil {
		return nil, delta, err
	}

	leftExpr := lang.NewTermExpr(ctx, term)

	if delta.token.IsOperator() && (delta.token.Is(syntax.Sum) || delta.token.Is(syntax.Minus)) {
		return c.exprAux(ctx, leftExpr, delta)
	}

	return leftExpr, delta, nil
}

func (c *Parser) exprAux(ctx *context.Context, left lang.Expr, delta TokenDelta) (lang.Expr, TokenDelta, error) {
	if !delta.token.IsOperator() {
		return lang.TermExpr{}, delta, unexpectedTokenTypeError(delta, tokens.Operator)
	}

	var leftExpr lang.Expr
	var err error

	if delta.token.Is(syntax.Sum) {
		leftExpr, delta, err = c.sumExpr(ctx, left, delta)
	} else if delta.token.Is(syntax.Minus) {
		leftExpr, delta, err = c.subtractExpr(ctx, left, delta)
	}

	if err != nil {
		return leftExpr, delta, err
	}

	if delta.token.IsOperator() && (delta.token.Is(syntax.Sum) || delta.token.Is(syntax.Minus)) {
		return c.exprAux(ctx, leftExpr, delta)
	}

	return leftExpr, delta, nil
}

func (c *Parser) sumExpr(ctx *context.Context, left lang.Expr, delta TokenDelta) (lang.SumExpr, TokenDelta, error) {
	if !delta.token.Is(syntax.Sum) {
		return lang.SumExpr{}, delta, unexpectedContentError(delta, syntax.Sum)
	}

	delta, err := c.nextToken()
	if err != nil {
		return lang.SumExpr{}, delta, err
	}

	factor, delta, err := c.Factor(ctx, delta)
	if err != nil {
		return lang.SumExpr{}, delta, err
	}

	term, delta, err := c.term(ctx, factor, delta)
	if err != nil {
		return lang.SumExpr{}, delta, err
	}

	sumExpr, err := lang.NewSumExpr(ctx, left, term)
	return sumExpr, delta, err
}

func (c *Parser) subtractExpr(ctx *context.Context, left lang.Expr, delta TokenDelta) (lang.SubtractExpr, TokenDelta, error) {
	if !delta.token.Is(syntax.Minus) {
		return lang.SubtractExpr{}, delta, unexpectedContentError(delta, syntax.Minus)
	}

	delta, err := c.nextToken()
	if err != nil {
		return lang.SubtractExpr{}, delta, err
	}

	factor, delta, err := c.Factor(ctx, delta)
	if err != nil {
		return lang.SubtractExpr{}, delta, err
	}

	term, delta, err := c.term(ctx, factor, delta)
	if err != nil {
		return lang.SubtractExpr{}, delta, err
	}

	delta, err = c.nextToken()
	if err != nil {
		return lang.SubtractExpr{}, delta, err
	}

	subExpr, err := lang.NewSubtractExpr(ctx, left, term)
	return subExpr, delta, err
}

func (c *Parser) term(ctx *context.Context, left lang.Factor, delta TokenDelta) (lang.Term, TokenDelta, error) {
	leftTerm := lang.NewFactorTerm(ctx, left)

	if delta.token.IsOperator() && (delta.token.Is(syntax.Multiply) || delta.token.Is(syntax.Divide)) {
		return c.termAux(ctx, leftTerm, delta)
	}

	return leftTerm, delta, nil
}

func (c *Parser) termAux(ctx *context.Context, left lang.Term, delta TokenDelta) (lang.Term, TokenDelta, error) {
	if !delta.token.IsOperator() {
		return lang.MultiplyTerm{}, delta, unexpectedTokenTypeError(delta, tokens.Operator)
	}

	var leftTerm lang.Term
	var err error

	if delta.token.Is(syntax.Multiply) {
		leftTerm, delta, err = c.multiplyTermAux(ctx, left, delta)
	} else if delta.token.Is(syntax.Divide) {
		leftTerm, delta, err = c.divideTermAux(ctx, left, delta)
	}

	if err != nil {
		return leftTerm, delta, err
	}

	if delta.token.IsOperator() && (delta.token.Is(syntax.Multiply) || delta.token.Is(syntax.Divide)) {
		return c.termAux(ctx, leftTerm, delta)
	}

	return leftTerm, delta, nil
}

func (c *Parser) multiplyTermAux(ctx *context.Context, left lang.Term, delta TokenDelta) (lang.MultiplyTerm, TokenDelta, error) {
	if !delta.token.Is(syntax.Multiply) {
		return lang.MultiplyTerm{}, delta, unexpectedContentError(delta, syntax.Multiply)
	}

	leftType, err := left.ResultingType()
	if err != nil {
		return lang.MultiplyTerm{}, delta, err
	}

	leftTypeMultipliable, ok := leftType.(value_types.MultipliableValueType)
	if !ok {
		return lang.MultiplyTerm{}, delta, failure.TypeNotMultipliable(leftType.Name())
	}

	delta, err = c.nextToken()
	if err != nil {
		return lang.MultiplyTerm{}, delta, err
	}

	factor, delta, err := c.Factor(ctx, delta)
	if err != nil {
		return lang.MultiplyTerm{}, delta, err
	}

	factorType, err := factor.ResultingType()
	if err != nil {
		return lang.MultiplyTerm{}, delta, err
	}

	_, err = leftTypeMultipliable.ResultingMultiplicationType(factorType)
	if err != nil {
		return lang.MultiplyTerm{}, delta, err
	}

	return lang.NewMultiplyTerm(ctx, left, factor), delta, err
}

func (c *Parser) divideTermAux(ctx *context.Context, left lang.Term, delta TokenDelta) (lang.DivideTerm, TokenDelta, error) {
	if !delta.token.Is(syntax.Divide) {
		return lang.DivideTerm{}, delta, unexpectedContentError(delta, syntax.Divide)
	}

	leftType, err := left.ResultingType()
	if err != nil {
		return lang.DivideTerm{}, delta, err
	}

	leftTypeMultipliable, ok := leftType.(value_types.DivisibleValueType)
	if !ok {
		return lang.DivideTerm{}, delta, failure.TypeNotDivisible(leftType.Name())
	}

	delta, err = c.nextToken()
	if err != nil {
		return lang.DivideTerm{}, delta, err
	}

	factor, delta, err := c.Factor(ctx, delta)
	if err != nil {
		return lang.DivideTerm{}, delta, err
	}

	factorType, err := factor.ResultingType()
	if err != nil {
		return lang.DivideTerm{}, delta, err
	}

	_, err = leftTypeMultipliable.ResultingDivisionType(factorType)
	if err != nil {
		return lang.DivideTerm{}, delta, err
	}

	return lang.NewDivideTerm(ctx, left, factor), delta, err
}

func (c *Parser) Factor(ctx *context.Context, delta TokenDelta) (lang.Factor, TokenDelta, error) {
	if delta.token.IsInteger() {
		return c.IntegerFactor(ctx, delta)
	}

	if delta.token.IsDecimal() {
		return c.DecimalFactor(ctx, delta)
	}

	if delta.token.IsIdentifier() {
		return c.SymbolFactor(ctx, delta)
	}

	if delta.token.IsString() {
		return c.StringFactor(ctx, delta)
	}

	if delta.token.IsOpenParenthesis() {
		return c.ExpressionFactor(ctx, delta)
	}

	return lang.ExpressionFactor{}, delta, noMatchTypeError(delta)
}

func (c *Parser) IntegerFactor(ctx *context.Context, delta TokenDelta) (lang.IntegerFactor, TokenDelta, error) {
	if !delta.token.IsInteger() {
		return lang.IntegerFactor{}, delta, unexpectedTokenTypeError(delta, tokens.Integer)
	}

	// -> 100
	value, err := strconv.Atoi(delta.token.Content())
	if err != nil {
		return lang.IntegerFactor{}, delta, err
	}

	integerFactor := lang.NewIntegerFactor(ctx, value)

	// deltas
	delta, err = c.nextToken()
	if err != nil {
		return integerFactor, delta, err
	}

	return integerFactor, delta, nil
}

func (c *Parser) DecimalFactor(ctx *context.Context, delta TokenDelta) (lang.FloatFactor, TokenDelta, error) {
	if !delta.token.IsDecimal() {
		return lang.FloatFactor{}, delta, unexpectedTokenTypeError(delta, tokens.Decimal)
	}

	// -> 100
	value, err := strconv.ParseFloat(delta.token.Content(), 64)
	if err != nil {
		return lang.FloatFactor{}, delta, err
	}

	floatFactor := lang.NewFloatFactor(ctx, value)

	// deltas
	delta, err = c.nextToken()
	if err != nil {
		return floatFactor, delta, err
	}

	return floatFactor, delta, nil
}

func (c *Parser) StringFactor(ctx *context.Context, delta TokenDelta) (lang.StringFactor, TokenDelta, error) {
	if !delta.token.IsString() {
		return lang.StringFactor{}, delta, unexpectedTokenTypeError(delta, tokens.String)
	}

	// -> "texto"
	stringFactor := lang.NewStringFactor(ctx, delta.token.Content())

	// deltas
	delta, err := c.nextToken()
	if err != nil {
		return stringFactor, delta, err
	}

	return stringFactor, delta, nil
}

func (c *Parser) SymbolFactor(ctx *context.Context, delta TokenDelta) (lang.SymbolFactor, TokenDelta, error) {
	if !delta.token.IsIdentifier() {
		return lang.SymbolFactor{}, delta, unexpectedTokenTypeError(delta, tokens.Identifier)
	}

	if !ctx.SymbolExists(delta.token.Content()) {
		return lang.SymbolFactor{}, delta, usedBeforeDeclaration(delta.token.Content())
	}

	symbol, err := ctx.RetrieveSymbol(delta.token.Content())
	if err != nil {
		return lang.SymbolFactor{}, delta, err
	}

	if !symbol.Assigned {
		return lang.SymbolFactor{}, delta, usedBeforeAssignment(symbol.Identifier)
	}

	symbolFactor := lang.NewSymbolFactor(ctx, symbol)

	delta, err = c.nextToken()
	if err != nil {
		return symbolFactor, delta, err
	}

	return symbolFactor, delta, nil
}

func (c *Parser) ExpressionFactor(ctx *context.Context, delta TokenDelta) (lang.ExpressionFactor, TokenDelta, error) {
	// -> (
	if !delta.token.IsOpenParenthesis() {
		return lang.ExpressionFactor{}, delta, unexpectedTokenTypeError(delta, tokens.OpenParenthesis)
	}

	delta, err := c.nextToken()
	if err != nil {
		return lang.ExpressionFactor{}, delta, err
	}

	// -> parte de dentro
	expr, delta, err := c.Expr(ctx, delta)
	if err != nil {
		return lang.ExpressionFactor{}, delta, err
	}

	expressionFactor := lang.NewExpressionFactor(ctx, expr)

	// -> )
	if !delta.token.IsCloseParenthesis() {
		return expressionFactor, delta, nil
	}

	delta, err = c.nextToken()
	if err != nil {
		return expressionFactor, delta, err
	}

	return expressionFactor, delta, nil
}
