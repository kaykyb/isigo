package lexer

import (
	"fmt"
	"isigo/common"
	"isigo/tokens"
)

type LexicalAnalysis struct {
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

func (l *LexicalAnalysis) SetContent(content string) error {
	if l.buffer != nil {
		return fmt.Errorf("Lexical analysis already contains a content buffer")
	}

	// Precisamos converter para rune para tratar caracteres UTF-8!
	l.buffer = []rune(content)
	l.bufferSize = len(l.buffer)

	return nil
}

func (l *LexicalAnalysis) NextToken() (token tokens.Token, tokenPosition common.CodePosition, err error) {
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
			err = &LexerError{
				err: err.Error(),
			}

			return token, l.position, err
		}
	}

	if tokenFound != nil {
		token = *tokenFound
	}

	return token, tokenPosition, err
}

func (l *LexicalAnalysis) Tokenize() (tokensFound []tokens.Token, err error) {
	token, _, err := l.NextToken()

	for token.Type() != tokens.EOF {
		tokensFound = append(tokensFound, token)
		token, _, err = l.NextToken()
	}

	return tokensFound, err
}

func (l *LexicalAnalysis) newPosition(delta ConsumptionDelta) common.CodePosition {
	newPosition := common.CodePosition{
		BufferPosition: l.position.BufferPosition,
		Line:           l.position.Line,
		Column:         l.position.Column,
	}

	if delta.linesConsumed > 0 {
		newPosition.Column = 0
	}

	newPosition.Column += delta.columnsConsumed
	newPosition.Line += delta.linesConsumed
	newPosition.BufferPosition += delta.runesConsumed

	return newPosition
}

func (l *LexicalAnalysis) decideNextTokenConsumer() (ConsumeTokenFunc, common.CodePosition, error) {
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
		return nil, l.position, &LexerError{
			err: fmt.Sprintf("'%c' não é um caractere permitido para iniciar um símbolo.", nextRune),
		}
	}
}

func (l *LexicalAnalysis) consumeWhitespaceRune() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	return delta, nil, nil
}

func (l *LexicalAnalysis) consumeNewlineRune() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		linesConsumed: 1,
		runesConsumed: 1,
	}

	return delta, nil, nil
}

func (l *LexicalAnalysis) consumeTabRune() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 4,
		runesConsumed:   1,
	}

	return delta, nil, nil
}

func (l *LexicalAnalysis) consumeCartridgeReturn() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		runesConsumed: 1,
	}

	return delta, nil, nil
}

func (l *LexicalAnalysis) consumeOperator() (ConsumptionDelta, *tokens.Token, error) {
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

func (l *LexicalAnalysis) consumeWord() (ConsumptionDelta, *tokens.Token, error) {
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

	token := tokens.NewIdentifier(word)

	if IsReservedWord(word) {
		token = tokens.NewReservedWord(word)
	}

	return delta, &token, nil
}

func (l *LexicalAnalysis) consumeNumber() (ConsumptionDelta, *tokens.Token, error) {
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

func (l *LexicalAnalysis) consumeAssignment() (ConsumptionDelta, *tokens.Token, error) {
	proceedingRune := l.peek(1)

	if proceedingRune != '=' {
		info := ConsumptionDelta{
			columnsConsumed: 2,
			runesConsumed:   2,
		}

		return info, nil, fmt.Errorf("Malformed assignment operator. Expected =, got %c", proceedingRune)
	}

	delta := ConsumptionDelta{
		columnsConsumed: 2,
		runesConsumed:   2,
	}

	token := tokens.NewAssign(":=")

	return delta, &token, nil
}

func (l *LexicalAnalysis) consumeStatementTerminator() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewStatementTerminator(".")

	return delta, &token, nil
}

func (l *LexicalAnalysis) consumeOpenParenthesis() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewOpenParenthesis("(")

	return delta, &token, nil
}

func (l *LexicalAnalysis) consumeCloseParenthesis() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewCloseParenthesis(")")

	return delta, &token, nil
}

func (l *LexicalAnalysis) consumeOpenBrace() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewOpenBrace("{")

	return delta, &token, nil
}

func (l *LexicalAnalysis) consumeCloseBrace() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewCloseBrace("}")

	return delta, &token, nil
}

func (l *LexicalAnalysis) consumeSeparator() (ConsumptionDelta, *tokens.Token, error) {
	delta := ConsumptionDelta{
		columnsConsumed: 1,
		runesConsumed:   1,
	}

	token := tokens.NewSeparator(",")

	return delta, &token, nil
}

func (l *LexicalAnalysis) consumeString() (ConsumptionDelta, *tokens.Token, error) {
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
		err = fmt.Errorf("Expected '\"' (quote) to conclude string of text")
	}

	return delta, &token, err
}

func (l *LexicalAnalysis) peek(d int) rune {
	if l.position.BufferPosition+d < l.bufferSize {
		return l.buffer[l.position.BufferPosition+d]
	}

	return 0
}
