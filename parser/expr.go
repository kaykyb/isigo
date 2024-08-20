package parser

import (
	"isigo/ast"
	"isigo/compiler_error"
	"isigo/context"
	"isigo/syntax"
	"isigo/tokens"
	"isigo/value_types"
	"strconv"
)

func (c *Parser) Expr(ctx *context.Context, delta TokenDelta) (ast.Expr, TokenDelta, error) {
	factor, delta, err := c.Factor(ctx, delta)
	if err != nil {
		return ast.TermExpr{}, delta, err
	}

	return c.exprInternal(ctx, factor, delta)
}

func (c *Parser) exprInternal(ctx *context.Context, left ast.Expr, delta TokenDelta) (ast.Expr, TokenDelta, error) {
	if delta.token.IsOperator() && (delta.token.Is(syntax.Sum) || delta.token.Is(syntax.Minus)) {
		return c.exprAux(ctx, left, delta)
	}

	return c.term(ctx, left, delta)
}

func (c *Parser) exprAux(ctx *context.Context, left ast.Expr, delta TokenDelta) (ast.Expr, TokenDelta, error) {
	if !delta.token.IsOperator() {
		return ast.TermExpr{}, delta, unexpectedTokenTypeError(delta, tokens.Operator)
	}

	var leftExpr ast.Expr
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

func (c *Parser) sumExpr(ctx *context.Context, left ast.Term, delta TokenDelta) (ast.SumExpr, TokenDelta, error) {
	if !delta.token.Is(syntax.Sum) {
		return ast.SumExpr{}, delta, unexpectedContentError(delta, syntax.Sum)
	}

	leftType, err := left.ResultingType()
	if err != nil {
		return ast.SumExpr{}, delta, err
	}

	leftTypeSumable, ok := leftType.(value_types.SumableValueType)
	if !ok {
		return ast.SumExpr{}, delta, compiler_error.TypeNotSumable(leftType.Name())
	}

	delta, err = c.nextToken()
	if err != nil {
		return ast.SumExpr{}, delta, err
	}

	factor, delta, err := c.Factor(ctx, delta)
	if err != nil {
		return ast.SumExpr{}, delta, err
	}

	term, delta, err := c.term(ctx, factor, delta)
	if err != nil {
		return ast.SumExpr{}, delta, err
	}

	termType, err := term.ResultingType()
	if err != nil {
		return ast.SumExpr{}, delta, err
	}

	_, err = leftTypeSumable.ResultingSumType(termType)
	if err != nil {
		return ast.SumExpr{}, delta, err
	}

	leftTerm := ast.NewTermExpr(ctx, left)
	return ast.NewSumExpr(ctx, leftTerm, term), delta, err
}

func (c *Parser) subtractExpr(ctx *context.Context, left ast.Term, delta TokenDelta) (ast.SubtractExpr, TokenDelta, error) {
	if !delta.token.Is(syntax.Minus) {
		return ast.SubtractExpr{}, delta, unexpectedContentError(delta, syntax.Minus)
	}

	leftType, err := left.ResultingType()
	if err != nil {
		return ast.SubtractExpr{}, delta, err
	}

	leftTypeSubtractable, ok := leftType.(value_types.SubtractableValueType)
	if !ok {
		return ast.SubtractExpr{}, delta, compiler_error.TypeNotSubtractable(leftType.Name())
	}

	delta, err = c.nextToken()
	if err != nil {
		return ast.SubtractExpr{}, delta, err
	}

	factor, delta, err := c.Factor(ctx, delta)
	if err != nil {
		return ast.SubtractExpr{}, delta, err
	}

	term, delta, err := c.term(ctx, factor, delta)
	if err != nil {
		return ast.SubtractExpr{}, delta, err
	}

	termType, err := term.ResultingType()
	if err != nil {
		return ast.SubtractExpr{}, delta, err
	}

	_, err = leftTypeSubtractable.ResultingSubtractType(termType)
	if err != nil {
		return ast.SubtractExpr{}, delta, err
	}

	delta, err = c.nextToken()
	if err != nil {
		return ast.SubtractExpr{}, delta, err
	}

	leftTerm := ast.NewTermExpr(ctx, left)
	return ast.NewSubtractExpr(ctx, leftTerm, term), delta, err
}

func (c *Parser) term(ctx *context.Context, left ast.Term, delta TokenDelta) (ast.Term, TokenDelta, error) {
	if delta.token.IsOperator() && (delta.token.Is(syntax.Multiply) || delta.token.Is(syntax.Divide)) {
		return c.termAux(ctx, left, delta)
	}

	return ast.NewFactorTerm(ctx, left), delta, nil
}

func (c *Parser) termAux(ctx *context.Context, left ast.Term, delta TokenDelta) (ast.Term, TokenDelta, error) {
	if !delta.token.IsOperator() {
		return ast.MultiplyTerm{}, delta, unexpectedTokenTypeError(delta, tokens.Operator)
	}

	var leftTerm ast.Term
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

func (c *Parser) multiplyTermAux(ctx *context.Context, left ast.Factor, delta TokenDelta) (ast.MultiplyTerm, TokenDelta, error) {
	if !delta.token.Is(syntax.Multiply) {
		return ast.MultiplyTerm{}, delta, unexpectedContentError(delta, syntax.Multiply)
	}

	leftType, err := left.ResultingType()
	if err != nil {
		return ast.MultiplyTerm{}, delta, err
	}

	leftTypeMultipliable, ok := leftType.(value_types.MultipliableValueType)
	if !ok {
		return ast.MultiplyTerm{}, delta, compiler_error.TypeNotMultipliable(leftType.Name())
	}

	delta, err = c.nextToken()
	if err != nil {
		return ast.MultiplyTerm{}, delta, err
	}

	factor, delta, err := c.Factor(ctx, delta)
	if err != nil {
		return ast.MultiplyTerm{}, delta, err
	}

	factorType, err := factor.ResultingType()
	if err != nil {
		return ast.MultiplyTerm{}, delta, err
	}

	_, err = leftTypeMultipliable.ResultingMultiplicationType(factorType)
	if err != nil {
		return ast.MultiplyTerm{}, delta, err
	}

	leftTerm := ast.NewFactorTerm(ctx, left)
	return ast.NewMultiplyTerm(ctx, leftTerm, factor), delta, err
}

func (c *Parser) divideTermAux(ctx *context.Context, left ast.Factor, delta TokenDelta) (ast.DivideTerm, TokenDelta, error) {
	if !delta.token.Is(syntax.Divide) {
		return ast.DivideTerm{}, delta, unexpectedContentError(delta, syntax.Divide)
	}

	leftType, err := left.ResultingType()
	if err != nil {
		return ast.DivideTerm{}, delta, err
	}

	leftTypeMultipliable, ok := leftType.(value_types.DivisibleValueType)
	if !ok {
		return ast.DivideTerm{}, delta, compiler_error.TypeNotDivisible(leftType.Name())
	}

	delta, err = c.nextToken()
	if err != nil {
		return ast.DivideTerm{}, delta, err
	}

	factor, delta, err := c.Factor(ctx, delta)
	if err != nil {
		return ast.DivideTerm{}, delta, err
	}

	factorType, err := factor.ResultingType()
	if err != nil {
		return ast.DivideTerm{}, delta, err
	}

	_, err = leftTypeMultipliable.ResultingDivisionType(factorType)
	if err != nil {
		return ast.DivideTerm{}, delta, err
	}

	leftTerm := ast.NewFactorTerm(ctx, left)
	return ast.NewDivideTerm(ctx, leftTerm, factor), delta, err
}

func (c *Parser) Factor(ctx *context.Context, delta TokenDelta) (ast.Factor, TokenDelta, error) {
	if delta.token.IsInteger() {
		return c.IntegerFactor(ctx, delta)
	}

	if delta.token.IsDecimal() {
		return c.DecimalFactor(ctx, delta)
	}

	if delta.token.IsIdentifier() {
		return c.SymbolFactor(ctx, delta)
	}

	if delta.token.IsOpenParenthesis() {
		return c.ExpressionFactor(ctx, delta)
	}

	return ast.ExpressionFactor{}, delta, noMatchTypeError(delta)
}

func (c *Parser) IntegerFactor(ctx *context.Context, delta TokenDelta) (ast.IntegerFactor, TokenDelta, error) {
	if !delta.token.IsInteger() {
		return ast.IntegerFactor{}, delta, unexpectedTokenTypeError(delta, tokens.Integer)
	}

	// -> 100
	value, err := strconv.Atoi(delta.token.Content())
	if err != nil {
		return ast.IntegerFactor{}, delta, err
	}

	integerFactor := ast.NewIntegerFactor(ctx, value)

	// deltas
	delta, err = c.nextToken()
	if err != nil {
		return integerFactor, delta, err
	}

	return integerFactor, delta, nil
}

func (c *Parser) DecimalFactor(ctx *context.Context, delta TokenDelta) (ast.FloatFactor, TokenDelta, error) {
	if !delta.token.IsDecimal() {
		return ast.FloatFactor{}, delta, unexpectedTokenTypeError(delta, tokens.Decimal)
	}

	// -> 100
	value, err := strconv.ParseFloat(delta.token.Content(), 64)
	if err != nil {
		return ast.FloatFactor{}, delta, err
	}

	floatFactor := ast.NewFloatFactor(ctx, value)

	// deltas
	delta, err = c.nextToken()
	if err != nil {
		return floatFactor, delta, err
	}

	return floatFactor, delta, nil
}

func (c *Parser) SymbolFactor(ctx *context.Context, delta TokenDelta) (ast.SymbolFactor, TokenDelta, error) {
	if !delta.token.IsIdentifier() {
		return ast.SymbolFactor{}, delta, unexpectedTokenTypeError(delta, tokens.Identifier)
	}

	if !ctx.SymbolExists(delta.token.Content()) {
		return ast.SymbolFactor{}, delta, usedBeforeDeclaration(delta.token.Content())
	}

	// TODO: Checa se foi inicializado

	symbol, err := ctx.RetrieveSymbol(delta.token.Content())
	if err != nil {
		return ast.SymbolFactor{}, delta, err
	}

	symbolFactor := ast.NewSymbolFactor(ctx, symbol)

	delta, err = c.nextToken()
	if err != nil {
		return symbolFactor, delta, err
	}

	return symbolFactor, delta, nil
}

func (c *Parser) ExpressionFactor(ctx *context.Context, delta TokenDelta) (ast.ExpressionFactor, TokenDelta, error) {
	// -> (
	if !delta.token.IsOpenParenthesis() {
		return ast.ExpressionFactor{}, delta, unexpectedTokenTypeError(delta, tokens.OpenParenthesis)
	}

	delta, err := c.nextToken()
	if err != nil {
		return ast.ExpressionFactor{}, delta, err
	}

	// -> parte de dentro
	expr, delta, err := c.Expr(ctx, delta)
	if err != nil {
		return ast.ExpressionFactor{}, delta, err
	}

	expressionFactor := ast.NewExpressionFactor(ctx, expr)

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
