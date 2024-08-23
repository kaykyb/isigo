package failure

type Environment int

const (
	LexerEnvironment Environment = iota
	SyntaxEnvironment
	SemanticEnvironment
)

type Failure interface {
	Environment() Environment
}
