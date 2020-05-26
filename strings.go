package types

import "strings"

// StringBuilder is an alias for strings.Builder
type StringBuilder = strings.Builder

// StringReader = strings.Reader
type StringReader = strings.Reader

// Replacer = strings.Replacer
type Replacer = strings.Replacer

// NewReader uses strings.NewReader
func NewReader(s string) *StringReader { return strings.NewReader(s) }

// NewReplacer returns strings.NewReplacer
func NewReplacer(oldnew ...string) *Replacer { return strings.NewReplacer(oldnew...) }
