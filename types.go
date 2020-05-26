// Package types is a type declaration and conversion package
package types

// var growFactor = 32 // tunes grow behavior

// poolStringBuilder is a global variable for pooling StringBuilder
//
// sb := poolStringBuilder.Get().(StringBuilder)
// ...
// sb.Reset()
// poolStringBuilder.Put(sb)
var poolStringBuilder = Pool{
	New: func() interface{} {
		return &StringBuilder{}
	},
}

func ceil32(int int) int {
	if int%32 != 0 {
		int += 32 - (int % 32)
	}
	return int
}
