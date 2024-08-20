package lexer

/*import (
	"reflect"
	"testing"
)

const testContent string = `Karma police, arrest this man
He talks in maths
He buzzes like a fridge
He's like a detuned radio

Karma police, arrest this girl
Her Hitler hairdo
Is making me feel ill
And we have crashed her party

This is what you'll get
This is what you'll get
This is what you'll get
When you mess with us

Karma police, I've given all I can
It's not enough
I've given all I can
But we're still on the payroll

This is what you'll get
This is what you'll get
This is what you'll get
When you mess with us

For a minute there
I lost myself, I lost myself
Phew, for a minute there
I lost myself, I lost myself

For a minute there
I lost myself, I lost myself
Phew, for a minute there
I lost myself, I lost myself`

func TestSetContentSetsBufferContent(t *testing.T) {
	l := LexicalAnalysis{}

	err := l.SetContent(testContent)

	if err != nil {
		t.Errorf(err.Error())
	}

	got := l.buffer
	want := []rune(testContent)

	if got == nil {
		t.Errorf("got %c, wanted %c", got, want)
	}
}

func TestSetContentSetsBufferContentOnlyOnce(t *testing.T) {
	l := LexicalAnalysis{}

	err := l.SetContent(testContent)

	if err != nil {
		t.Errorf(err.Error())
	}

	got := l.SetContent(testContent)

	if got == nil {
		t.Errorf("got %q, wanted an compiler_error", got)
	}
}

func TestTokenizeWhitespacesString(t *testing.T) {
	content := " \t\n  \r\n  "

	l := LexicalAnalysis{}
	err := l.SetContent(content)

	if err != nil {
		t.Errorf(err.Error())
	}

	got, err := l.Tokenize()

	if err != nil {
		t.Errorf(err.Error())
	}

	wants := []Token{}

	if len(got) != 0 {
		t.Errorf("Expected slices to be equal: %v and %v", got, wants)
	}
}

func TestTokenizeOperatorsString(t *testing.T) {
	content := " +-*/ /* > < >= <= != == <+>+>"

	l := LexicalAnalysis{}
	err := l.SetContent(content)

	if err != nil {
		t.Errorf(err.Error())
	}

	got, err := l.Tokenize()

	if err != nil {
		t.Errorf(err.Error())
	}

	wants := []Token{
		{internalType: OPERATOR, content: "+"},
		{internalType: OPERATOR, content: "-"},
		{internalType: OPERATOR, content: "*"},
		{internalType: OPERATOR, content: "/"},
		{internalType: OPERATOR, content: ">"},
		{internalType: OPERATOR, content: "<"},
		{internalType: OPERATOR, content: ">="},
		{internalType: OPERATOR, content: "<="},
		{internalType: OPERATOR, content: "!="},
		{internalType: OPERATOR, content: "=="},
		{internalType: OPERATOR, content: "<"},
		{internalType: OPERATOR, content: "+"},
		{internalType: OPERATOR, content: ">"},
		{internalType: OPERATOR, content: "+"},
		{internalType: OPERATOR, content: ">"},
	}

	if !reflect.DeepEqual(got, wants) {
		t.Errorf("Expected slices to be equal: %v and %v", got, wants)
	}
}

func TestTokenizeNumbersString(t *testing.T) {
	content := "1 1.0 2.024 .1 2223.3 2222"

	l := LexicalAnalysis{}
	err := l.SetContent(content)

	if err != nil {
		t.Errorf(err.Error())
	}

	got, err := l.Tokenize()

	if err != nil {
		t.Errorf(err.Error())
	}

	wants := []Token{
		{internalType: INTEGER, content: "1"},
		{internalType: DECIMAL, content: "1.0"},
		{internalType: DECIMAL, content: "2.024"},
		{internalType: DECIMAL, content: ".1"},
		{internalType: DECIMAL, content: "2223.3"},
		{internalType: INTEGER, content: "2222"},
	}

	if !reflect.DeepEqual(got, wants) {
		t.Errorf("Expected slices to be equal: %v and %v", got, wants)
	}
}

func TestTokenizeStringString(t *testing.T) {
	content := "1 \"abc def 83has901 asduh aac cccc\""

	l := LexicalAnalysis{}
	err := l.SetContent(content)

	if err != nil {
		t.Errorf(err.Error())
	}

	got, err := l.Tokenize()

	if err != nil {
		t.Errorf(err.Error())
	}

	wants := []Token{
		{internalType: INTEGER, content: "1"},
		{internalType: STRING, content: "abc def 83has901 asduh aac cccc"},
	}

	if !reflect.DeepEqual(got, wants) {
		t.Errorf("Expected slices to be equal: %v and %v", got, wants)
	}
}*/
