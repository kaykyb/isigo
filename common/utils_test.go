package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndentBasic(t *testing.T) {
	input := "line1\nline2\nline3"
	expected := "\tline1\n\tline2\n\tline3"
	result := Indent(input)

	assert.Equal(t, expected, result, "Indent should indent each line with a tab.")
}

func TestIndentSingleLine(t *testing.T) {
	input := "singleline"
	expected := "\tsingleline"
	result := Indent(input)

	assert.Equal(t, expected, result, "Indent should indent a single line with a tab.")
}

func TestIndentEmptyString(t *testing.T) {
	input := ""
	expected := "\t"
	result := Indent(input)

	assert.Equal(t, expected, result, "Indent should return an empty string when input is empty.")
}

func TestIndentNewLinesOnly(t *testing.T) {
	input := "\n\n\n"
	expected := "\t\n\t\n\t\n\t"
	result := Indent(input)

	assert.Equal(t, expected, result, "Indent should handle strings with only new lines correctly.")
}

func TestIndentMixedContent(t *testing.T) {
	input := "line1\n\nline2\n\n"
	expected := "\tline1\n\t\n\tline2\n\t\n\t"
	result := Indent(input)

	assert.Equal(t, expected, result, "Indent should handle mixed content correctly.")
}
