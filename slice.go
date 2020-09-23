package types

import "sort"

// Slice is a slice of any
type Slice = []I

// SliceEncoder writes Slice to string
type SliceEncoder = func(Slice) string

// SliceInt = sort.IntSlice
type SliceInt = sort.IntSlice

// SliceReflect returns a Slice from an arg
func SliceReflect(arg I) (out Slice) {
	if val := Reflect(arg); val.Kind() == KindSlice {
		out = sliceValue(val)
	}
	return
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

// SliceUI is []uint with sort-like additions
type SliceUI []uint

func (p SliceUI) Len() int           { return len(p) }
func (p SliceUI) Less(i, j int) bool { return p[i] < p[j] }
func (p SliceUI) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Sort calls sort.Sort with this
func (p SliceUI) Sort() { sort.Sort(p) }
