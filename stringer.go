package types

import "fmt"

// Stringer = fmt.Stringer
type Stringer = fmt.Stringer

// NewStringer returns a wrapper Stringer that defers parsing the interal string
func NewStringer(any I) Stringer {
	return stringer{any}
}

type stringer struct {
	i I
}

func (i stringer) String() string {
	return NewString(i.i)
}
