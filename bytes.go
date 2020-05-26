package types

import "bytes"

// Bytes is a []byte
type Bytes = []byte

// Buffer = bytes.Buffer
type Buffer = bytes.Buffer

// NewBuffer returns Bytes.NewBuffer
func NewBuffer(b Bytes) *Buffer {
	return bytes.NewBuffer(b)
}

// NewBufferString returns bytes.NewBufferString
func NewBufferString(s string) *Buffer {
	return bytes.NewBufferString(s)
}

// BytesS casts string to []byte
//
// outperforms `[]byte(s)` by ~95%
func BytesS(s string) []byte {
	return *(*[]byte)(Pointer(&s))
}
