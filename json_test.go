package types_test

import (
	"testing"

	"taylz.io/types"
)

func TestJSONString(t *testing.T) {
	json := types.Dict{
		"hello": types.NewStringer("world"),
	}
	jsons := types.NewStringDict(json)
	if jsons != `{"hello":world}` {
		t.Fail()
	}
}

func TestEncodeJSON(t *testing.T) {
	ans := types.Slice{"hello world", false, map[int]bool{20: true}, map[string]string{"x": "19"}, types.NewStringer(`foobar`)}
	anss := `["hello world",false,{20:true},{"x":"19"},foobar]`
	if str := types.NewStringSlice(ans); str != anss {
		t.Log(`Expected ` + anss)
		t.Log(`Received ` + str)
		t.Fail()
	}
}
