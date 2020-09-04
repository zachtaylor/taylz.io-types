package types

// Err creates an error
func Err(str string) error {
	return Error{
		Text: str,
	}
}

// ErrWith creates an error with a source
func ErrWith(str string, source error) error {
	return Error{
		Text:   str,
		Source: source,
	}
}

// Error is a basic error
type Error struct {
	Text   string
	Source error
}

func (e Error) Error() string {
	s := e.Text
	if e.Source != nil {
		s += ": " + e.Source.Error()
	}
	return s
}

func (e Error) String() string { return "Error{" + e.Error() + "}" }
