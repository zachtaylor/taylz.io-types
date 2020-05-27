package types

import (
	"sort"
	"strings"
)

// Slice is a slice of any
type Slice = []I

// SliceEncoder writes Slice to string
type SliceEncoder = func(Slice) string

// SliceInt = sort.IntSlice
type SliceInt = sort.IntSlice

// SliceReflect returns a Slice from an arg
func SliceReflect(arg I) (out Slice) {
	if val := Reflect(arg); val.Kind() == SliceKind {
		out = sliceValue(val)
	}
	return
}

// SliceSort uses sort.Slice
func SliceSort(slice Slice, less func(i int, j int) bool) {
	sort.Slice(slice, less)
}

func sliceValue(val Value) (out Slice) {
	len := val.Len()
	out = make(Slice, len)
	for i := 0; i < len; i++ {
		out[i] = val.Index(i).Interface()
	}
	return
}

// SliceString = sort.StringSlice
type SliceString = sort.StringSlice

// NewSliceStringSplit is an alias for strings.Split
func NewSliceStringSplit(s, sep string) SliceString { return strings.Split(s, sep) }
