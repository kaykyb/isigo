package sources

type SourceStream interface {
	Peek(d int) rune
	FlushRune() error
	FlushMultipleRunes(runes int) error
}
