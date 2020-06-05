package types

import "fmt"

// Stringer = fmt.Stringer
type Stringer = fmt.Stringer

// NewStringer returns a wrapper Stringer that defers parsing the interal string
func NewStringer(i I) Stringer { return stringer{i} }

type stringer struct{ I } // generic container

func (stringer stringer) String() string { return NewString(stringer.I) }

// StringerDict is a Stringer wrapper for Dict
type StringerDict Dict

func (dict StringerDict) String() string { return NewStringDict(dict) }

// StringerMap is a Stringer wrapper for Map
type StringerMap Map

func (m StringerMap) String() string { return NewStringMap(m) }

// StringerSlice is a Stringer wrapper for Slice
type StringerSlice Slice

func (slice StringerSlice) String() string { return NewStringSlice(slice) }
