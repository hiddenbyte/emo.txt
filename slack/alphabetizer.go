package slack

import (
	"bufio"
	"fmt"
	"io"

	"github.com/hiddenbyte/emo.txt/core"
)

type alphabetizer struct {
	lastRuneWasAlha bool
	w               *bufio.Writer
}

func (a *alphabetizer) WriteRune(r rune) error {
	if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
		if !a.lastRuneWasAlha {
			a.w.WriteString(fmt.Sprintf(":alphabet-yellow-%c:", r))
			a.lastRuneWasAlha = true
		} else {
			a.w.WriteString(fmt.Sprintf(":alphabet-white-%c:", r))
		}
	} else {
		a.lastRuneWasAlha = false
		a.w.WriteRune(r)
	}
	return nil
}

func (a *alphabetizer) Flush() error {
	return a.w.Flush()
}

// NewAlphabetizer creates a new alphabetizer
func NewAlphabetizer(w io.Writer) core.EmoWriter {
	return &alphabetizer{false, bufio.NewWriter(w)}
}
