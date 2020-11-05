package types_test

import (
	"strings"
	"testing"

	"taylz.io/types"
)

func TestJSONString(t *testing.T) {
	json := MakeDataDict1()
	jsons := types.StringDict(json)
	if strings.Contains(jsons, "\n") {
		t.Log("DataDict1 generated newline")
		t.Fail()
	}

	json = MakeDataDict2()
	jsons = types.StringDict(json)
	if strings.Contains(jsons, "\n") {
		t.Log("DataDict2 generated newline")
		t.Fail()
	}

	json = MakeDataDict3()
	jsons = types.StringDict(json)
	if strings.Contains(jsons, "\n") {
		t.Log("DataDict3 generated newline")
		t.Fail()
	}
}

func TestEncodeJSON(t *testing.T) {
	ans := types.Slice{"hello world", false, map[int]bool{20: true}, map[string]string{"x": "19"}}
	anss := `["hello world",false,{20:true},{"x":"19"}]`
	if str := types.StringSlice(ans); str != anss {
		t.Log(`Expected ` + anss)
		t.Log(`Received ` + str)
		t.Fail()
	}
}
