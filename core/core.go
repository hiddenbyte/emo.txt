package core

// EmoWriter emo writer
type EmoWriter interface {
	WriteRune(r rune) error
	Flush() error
}
