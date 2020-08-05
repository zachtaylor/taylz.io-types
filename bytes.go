package types

import "bytes"

// Buffer = bytes.Buffer
type Buffer = bytes.Buffer

// Bytes is a []byte
type Bytes = []byte

// NewBuffer returns Bytes.NewBuffer
func NewBuffer(b Bytes) *Buffer { return bytes.NewBuffer(b) }

// NewBufferString returns bytes.NewBufferString
func NewBufferString(s string) *Buffer { return bytes.NewBufferString(s) }

// BytesString casts string to Bytes
//
// outperforms `[]byte(s)` by ~95%
func BytesString(s string) Bytes { return *(*Bytes)(Pointer(&s)) }
