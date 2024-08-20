package lang

type Node interface {
	Output() (string, error)
}
