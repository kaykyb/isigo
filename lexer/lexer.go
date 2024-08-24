package lexer

import (
	"isigo/common"
	"isigo/failure"
	"isigo/tokens"
)

type Lexer struct {
	position   common.CodePosition
	buffer     []rune
	bufferSize int
}

type ConsumptionDelta struct {
	runesConsumed   int
	linesConsumed   int
	columnsConsumed int
}

type ConsumeTokenFunc = func() (delta ConsumptionDelta, token *tokens.Token, err error)

func New(content string) Lexer {
	runeBuffer := []rune(content)
	return Lexer{
		buffer:     []rune(content),
		bufferSize: len(runeBuffer),
	}
}

func (l *Lexer) NextToken() (token tokens.Token, tokenPosition common.CodePosition, err error) {
	tokenPosition = l.position

	var tokenFound *tokens.Token

	for l.position.BufferPosition < l.bufferSize && tokenFound == nil {
		var consume ConsumeTokenFunc
		consume, tokenPosition, err = l.decideNextTokenConsumer()

		if err != nil || consume == nil {
			return tokens.Token{}, tokenPosition, err
		}

		var delta ConsumptionDelta
		delta, tokenFound, err = consume()

		l.position = l.newPosition(delta)

		if err != nil {
			return token, l.position, err
		}
	}

	if tokenFound != nil {
		token = *tokenFound
	}

	return token, tokenPosition, err
}

func (l *Lexer) newPosition(delta ConsumptionDelta) common.CodePosition {
	newPosition := common.NewCodePosition(l.position.BufferPosition, l.position.Line, l.position.Column)

	if delta.linesConsumed > 0 {
		newPosition.Column = 0
	}

	newPosition.Column += delta.columnsConsumed
	newPosition.Line += delta.linesConsumed
	newPosition.BufferPosition += delta.runesConsumed

	return newPosition
}

func (l *Lexer) decideNextTokenConsumer() (ConsumeTokenFunc, common.CodePosition, error) {
	nextRune := l.peek(0)

	switch {
	case IsWhitespaceRune(nextRune):
		return l.consumeWhitespaceRune, l.position, nil
	case IsNewLineRune(nextRune):
		return l.consumeNewlineRune, l.position, nil
	case IsTabRune(nextRune):
		return l.consumeTabRune, l.position, nil
	case IsCartridgeReturnRune(nextRune):
		return l.consumeCartridgeReturn, l.position, nil
	case IsOperatorRune(nextRune):
		return l.consumeOperator, l.position, nil
	case IsLetterRune(nextRune):
		return l.consumeWord, l.position, nil
	case IsDigitRune(nextRune):
		return l.consumeNumber, l.position, nil
	case IsStatementTerminator(nextRune):
		return l.consumeStatementTerminator, l.position, nil
	case IsDecimalSeparator(nextRune):
		return l.consumeNumber, l.position, nil
	case IsColon(nextRune):
		return l.consumeAssignment, l.position, nil
	case IsOpenParenthesis(nextRune):
		return l.consumeOpenParenthesis, l.position, nil
	case IsCloseParenthesis(nextRune):
		return l.consumeCloseParenthesis, l.position, nil
	case IsOpenBrace(nextRune):
		return l.consumeOpenBrace, l.position, nil
	case IsCloseBrace(nextRune):
		return l.consumeCloseBrace, l.position, nil
	case IsComma(nextRune):
		return l.consumeSeparator, l.position, nil
	case IsStringDelimiter(nextRune):
		return l.consumeString, l.position, nil
	default:
		return nil, l.position, failure.UnexpectedCharacter(nextRune)
	}
}

func (l *Lexer) consumeWhitespaceRune() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	return delta, nil, nil
}

func (l *Lexer) consumeNewlineRune() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		linesConsumed: 1,
		runesConsumed: 1,
	}

	return delta, nil, nil
}

func (l *Lexer) consumeTabRune() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 4,
		runesConsumed:   1,
	}

	return delta, nil, nil
}

func (l *Lexer) consumeCartridgeReturn() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		runesConsumed: 1,
	}

	return delta, nil, nil
}

func (l *Lexer) consumeOperator() (ConsumptionDelta, *tokens.Token, error) {
	startRune := l.peek(0)
	proceedingRune := l.peek(1)

	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewOperator(string(startRune))

	if proceedingRune == '=' {
		delta.columnsConsumed = 2
		delta.runesConsumed = 2

		token = tokens.NewOperator(string(startRune) + string(proceedingRune))
	}

	return delta, &token, nil
}

func (l *Lexer) consumeWord() (ConsumptionDelta, *tokens.Token, error) {
	word := ""
	wDelta := 0

	for wDelta = 0; true; wDelta++ {
		nextRune := l.peek(wDelta)

		if !IsLetterRune(nextRune) && !IsDigitRune(nextRune) {
			break
		}

		word += string(nextRune)
	}

	delta := ConsumptionDelta{
		columnsConsumed: wDelta,
		runesConsumed:   wDelta,
	}

	var token tokens.Token

	if IsReservedWord(word) {
		token = tokens.NewReservedWord(word)
	} else if IsTypeT(word) {
		token = tokens.NewTypeT(word)
	} else {
		token = tokens.NewIdentifier(word)
	}

	return delta, &token, nil
}

func (l *Lexer) consumeNumber() (ConsumptionDelta, *tokens.Token, error) {
	word := ""
	isDecimal := false

	wDelta := 0

	for wDelta = 0; true; wDelta++ {
		nextRune := l.peek(wDelta)
		isDecimalSeparator := IsDecimalSeparator(nextRune)

		if !IsDigitRune(nextRune) && !isDecimalSeparator {
			break
		}

		// Já é decimal, apareceu um segundo ponto decimal
		if isDecimalSeparator && isDecimal {
			break
		}

		// Apareceu um ponto decimal nesta posição, mas o próximo runa não é um
		// digito
		if isDecimalSeparator && !IsDigitRune(l.peek(wDelta+1)) {
			break
		}

		if isDecimalSeparator {
			isDecimal = true
		}

		word += string(nextRune)
	}

	delta := ConsumptionDelta{
		columnsConsumed: wDelta,
		runesConsumed:   wDelta,
	}

	token := tokens.NewInteger(word)

	if isDecimal {
		token = tokens.NewDecimal(word)
	}

	return delta, &token, nil
}

func (l *Lexer) consumeAssignment() (ConsumptionDelta, *tokens.Token, error) {
	proceedingRune := l.peek(1)

	if proceedingRune != '=' {
		info := ConsumptionDelta{
			columnsConsumed: 2,
			runesConsumed:   2,
		}

		return info, nil, failure.MalformedAssignmentOperator(proceedingRune)
	}

	delta := ConsumptionDelta{
		columnsConsumed: 2,
		runesConsumed:   2,
	}

	token := tokens.NewAssign(":=")

	return delta, &token, nil
}

func (l *Lexer) consumeStatementTerminator() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewStatementTerminator(".")

	return delta, &token, nil
}

func (l *Lexer) consumeOpenParenthesis() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewOpenParenthesis("(")

	return delta, &token, nil
}

func (l *Lexer) consumeCloseParenthesis() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewCloseParenthesis(")")

	return delta, &token, nil
}

func (l *Lexer) consumeOpenBrace() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewOpenBrace("{")

	return delta, &token, nil
}

func (l *Lexer) consumeCloseBrace() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewCloseBrace("}")

	return delta, &token, nil
}

func (l *Lexer) consumeSeparator() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewSeparator(",")

	return delta, &token, nil
}

func (l *Lexer) consumeString() (ConsumptionDelta, *tokens.Token, error) {
	content := ""
	wDelta := 0

	isInText := false

	var err error

	for wDelta = 0; true; wDelta++ {
		nextRune := l.peek(wDelta)

		if IsNewLineRune(nextRune) || nextRune == 0 {
			break
		}

		if IsStringDelimiter(nextRune) && isInText {
			isInText = false
			wDelta++
			break
		}

		if IsStringDelimiter(nextRune) {
			isInText = true
		} else {
			content += string(nextRune)
		}
	}

	delta := ConsumptionDelta{
		columnsConsumed: wDelta,
		runesConsumed:   wDelta,
	}

	token := tokens.NewString(content)

	if isInText {
		err = failure.ExpectedEndQuote()
	}

	return delta, &token, err
}

func (l *Lexer) peek(d int) rune {
	if l.position.BufferPosition+d < l.bufferSize {
		return l.buffer[l.position.BufferPosition+d]
	}

	return 0
}
