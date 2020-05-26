package types

// Error is a basic error
type Error struct {
	Text   string
	Source error
}

func (e Error) isError() error {
	return e
}

func (e Error) Error() string {
	return e.Text + ": " + e.Source.Error()
}

func (e Error) String() string {
	return "Error{" + e.Text + " " + e.Source.Error() + "}"
}
