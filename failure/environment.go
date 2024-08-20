package failure

import "fmt"

type Environment int

const (
	LexerEnvironment Environment = iota
	SyntaxEnvironment
	SemanticEnvironment
)

type Failure interface {
	Environment() Environment
}

type Lexer struct {
	err string
}
type Syntax struct {
	err string
}
type Semantic struct {
	err string
}

func (e Lexer) Error() string {
	return fmt.Sprintf("[Lexical Error]: %s", e.err)
}

func (e Lexer) Environment() Environment {
	return LexerEnvironment
}

func LexerErrorf(format string, args ...interface{}) error {
	return Lexer{
		err: fmt.Sprintf(format, args...),
	}
}

func (e Syntax) Error() string {
	return fmt.Sprintf("[Syntax Error]: %s", e.err)
}

func (e Syntax) Environment() Environment {
	return SyntaxEnvironment
}

func SyntaxErrorf(format string, args ...interface{}) error {
	return Syntax{
		err: fmt.Sprintf(format, args...),
	}
}

func (e Semantic) Error() string {
	return fmt.Sprintf("[Semantic Error]: %s", e.err)
}

func (e Semantic) Environment() Environment {
	return SemanticEnvironment
}

func SemanticErrorf(format string, args ...interface{}) error {
	return Semantic{
		err: fmt.Sprintf(format, args...),
	}
}
