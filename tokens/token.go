package tokens

type Type int

const (
	EOF Type = iota

	Assign
	Operator

	OpenParenthesis
	CloseParenthesis
	OpenBrace
	CloseBrace

	StatementTerminator
	Separator

	Integer
	Decimal
	String

	Identifier
	ReservedWord
	TypeT
)

type Token struct {
	internalType Type
	content      string
}

func (t *Token) Type() Type {
	return t.internalType
}

func (t *Token) Content() string {
	return t.content
}

func (t *Token) FriendlyString() string {
	return FriendlyString(t.internalType)
}

func FriendlyString(ttype Type) string {
	friendlyStringMap := map[Type]string{
		EOF: "fim de arquivo",

		Assign:   "atribuição",
		Operator: "operador",

		OpenParenthesis:  "parentêses",
		CloseParenthesis: "parentêses",
		OpenBrace:        "chaves",
		CloseBrace:       "chaves",

		StatementTerminator: "fim de sentença",
		Separator:           "separador",

		Integer: "inteiro",
		Decimal: "decimal",
		String:  "cadeia de caracteres",

		Identifier:   "identificador",
		ReservedWord: "palavra reserva",
		TypeT:        "tipo",
	}

	return friendlyStringMap[ttype]
}

func newToken(ttype Type, content string) Token {
	return Token{
		internalType: ttype,
		content:      content,
	}
}

func (t *Token) Is(content string) bool {
	return t.content == content
}

func (t *Token) IsReservedWord() bool {
	return t.internalType == ReservedWord
}

func (t *Token) IsTypeT() bool {
	return t.internalType == TypeT
}

func (t *Token) IsIdentifier() bool {
	return t.internalType == Identifier
}

func (t *Token) IsAssign() bool {
	return t.internalType == Assign
}

func (t *Token) IsInteger() bool {
	return t.internalType == Integer
}

func (t *Token) IsDecimal() bool {
	return t.internalType == Decimal
}

func (t *Token) IsOperator() bool {
	return t.internalType == Operator
}

func (t *Token) IsSeparator() bool {
	return t.internalType == Separator
}

func (t *Token) IsOpenParenthesis() bool {
	return t.internalType == OpenParenthesis
}

func (t *Token) IsCloseParenthesis() bool {
	return t.internalType == CloseParenthesis
}

func (t *Token) IsOpenBrace() bool {
	return t.internalType == OpenBrace
}

func (t *Token) IsCloseBrace() bool {
	return t.internalType == CloseBrace
}

func (t *Token) IsStatementTerminator() bool {
	return t.internalType == StatementTerminator
}
