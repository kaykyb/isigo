package parser

import (
	"isigo/lang"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExprSum(t *testing.T) {
	p, c := SetupLPC(t, "1 + 2")
	delta := AssertNextToken(t, p)

	doc, delta, err := p.Expr(c, delta)
	assert.NoError(t, err)

	left := lang.NewTermExpr(c, lang.NewFactorTerm(c, lang.NewIntegerFactor(c, 1)))
	right := lang.NewFactorTerm(c, lang.NewIntegerFactor(c, 2))
	expectedExpr, _ := lang.NewSumExpr(c, left, right)

	assert.Equal(t, expectedExpr, doc)
}

func TestExprSubtract(t *testing.T) {
	p, c := SetupLPC(t, "1 - 2")
	delta := AssertNextToken(t, p)

	doc, delta, err := p.Expr(c, delta)
	assert.NoError(t, err)

	left := lang.NewTermExpr(c, lang.NewFactorTerm(c, lang.NewIntegerFactor(c, 1)))
	right := lang.NewFactorTerm(c, lang.NewIntegerFactor(c, 2))
	expectedExpr, _ := lang.NewSubtractExpr(c, left, right)

	assert.Equal(t, expectedExpr, doc)
}

func TestExprMultiply(t *testing.T) {
	p, c := SetupLPC(t, "1 * 2")
	delta := AssertNextToken(t, p)

	doc, delta, err := p.Expr(c, delta)
	assert.NoError(t, err)

	left := lang.NewFactorTerm(c, lang.NewIntegerFactor(c, 1))
	right := lang.NewIntegerFactor(c, 2)
	term := lang.NewMultiplyTerm(c, left, right)
	expectedExpr := lang.NewTermExpr(c, term)

	assert.Equal(t, expectedExpr, doc)
}

func TestExprDivide(t *testing.T) {
	p, c := SetupLPC(t, "1 / 2")
	delta := AssertNextToken(t, p)

	doc, delta, err := p.Expr(c, delta)
	assert.NoError(t, err)

	left := lang.NewFactorTerm(c, lang.NewIntegerFactor(c, 1))
	right := lang.NewIntegerFactor(c, 2)
	term := lang.NewDivideTerm(c, left, right)
	expectedExpr := lang.NewTermExpr(c, term)

	assert.Equal(t, expectedExpr, doc)
}

func TestExprOrders(t *testing.T) {
	p, c := SetupLPC(t, "8 / 2 * (2 + 2) + 1")
	delta := AssertNextToken(t, p)

	doc, delta, err := p.Expr(c, delta)
	assert.NoError(t, err)

	insideExpr, _ := lang.NewSumExpr(c,
		lang.NewTermExpr(c, lang.NewFactorTerm(c, lang.NewIntegerFactor(c, 2))),
		lang.NewFactorTerm(c, lang.NewIntegerFactor(c, 2)),
	)

	expectedExpr, _ :=
		lang.NewSumExpr(c,
			lang.NewTermExpr(
				c, lang.NewMultiplyTerm(c,
					lang.NewDivideTerm(c,
						lang.NewFactorTerm(c, lang.NewIntegerFactor(c, 8)),
						lang.NewIntegerFactor(c, 2),
					),
					lang.NewExpressionFactor(c,
						insideExpr,
					),
				),
			),
			lang.NewFactorTerm(c, lang.NewIntegerFactor(c, 1)),
		)

	assert.Equal(t, expectedExpr, doc)
}
