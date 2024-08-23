package common

import (
	"strings"
)

func Indent(input string) string {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		lines[i] = "\t" + line
	}

	return strings.Join(lines, "\n")
}
