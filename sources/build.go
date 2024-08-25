package sources

import (
	"bufio"
	"io"
)

type BuildReader struct {
	reader                 *bufio.Reader
	runesReadInCurrentLine int
	currentLine            string
	ended                  bool
}

func NewBuildReader(reader *bufio.Reader) *BuildReader {
	r := &BuildReader{
		reader:                 reader,
		runesReadInCurrentLine: 0,
		currentLine:            "",
		ended:                  false,
	}

	r.readNextLine()

	return r
}

func (r *BuildReader) FlushRune() error {
	r.runesReadInCurrentLine++

	if r.runesReadInCurrentLine >= len(r.currentLine) {
		r.runesReadInCurrentLine = 0
		return r.readNextLine()
	}

	return nil
}

func (r *BuildReader) FlushMultipleRunes(runes int) error {
	for i := 0; i < runes; i++ {
		if err := r.FlushRune(); err != nil {
			return err
		}
	}

	return nil
}

func (r *BuildReader) readNextLine() error {
	r.runesReadInCurrentLine = 0
	currentLine, err := r.reader.ReadString('\n')
	r.currentLine = currentLine

	if err == io.EOF {
		r.ended = true
	} else if err != nil {
		return err
	}

	return nil
}

func (r *BuildReader) Peek(d int) rune {
	if r.runesReadInCurrentLine+d >= len(r.currentLine) {
		return 0
	}

	return rune(r.currentLine[r.runesReadInCurrentLine+d])
}
