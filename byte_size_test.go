package types_test

import (
	"testing"

	"taylz.io/types"
)

func TestByteSizeDict(t *testing.T) {
	dict := types.Dict{}
	dict["hello"] = "world"
	size := len(types.StringDict(dict))
	if actualSize := types.ByteSizeJSONDict(dict); actualSize != size {
		t.Log("Expected", size)
		t.Log("Actual", actualSize)
		t.Fail()
	}
}

func TestByteSizeSlice(t *testing.T) {
	slice := types.Slice{}
	slice = append(slice, "hello", "world")
	size := len(types.StringSlice(slice))
	if actualSize := types.ByteSizeJSONSlice(slice); actualSize != size {
		t.Log("Expected", size)
		t.Log("Actual", actualSize)
		t.Fail()
	}
}
