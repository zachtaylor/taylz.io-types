package types_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"taylz.io/types"
)

func BenchmarkBuiltinStringBytes(b *testing.B) {
	x := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	for i := 0; i < b.N; i++ {
		_ = string(x)
	}
}

func BenchmarkCastStringBytes(b *testing.B) {
	x := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	for i := 0; i < b.N; i++ {
		_ = types.StringBytes(x)
	}
}

func BenchmarkSTDFmtSprintSlice(b *testing.B) {
	data := MakeDataSlice()
	for i := 0; i < b.N; i++ {
		fmt.Sprint(data)
	}
}

func BenchmarkSTDJsonMarshalSlice(b *testing.B) {
	data := MakeDataSlice()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

func BenchmarkTypesStringSlice(b *testing.B) {
	data := MakeDataSlice()
	for i := 0; i < b.N; i++ {
		types.String(data)
	}
}

func BenchmarkFmtSprint1(b *testing.B) {
	data := MakeDataDict1()
	for i := 0; i < b.N; i++ {
		fmt.Sprint(data)
	}
}

func BenchmarkJsonMarshal1(b *testing.B) {
	data := MakeDataDict1()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

func BenchmarkTypesDictString1(b *testing.B) {
	data := MakeDataDict1()
	for i := 0; i < b.N; i++ {
		types.String(data)
	}
}

func BenchmarkTypesMapString1(b *testing.B) {
	data := types.MapReflect(MakeDataDict1())
	for i := 0; i < b.N; i++ {
		types.String(data)
	}
}

func BenchmarkFmtSprint2(b *testing.B) {
	data := MakeDataDict2()
	for i := 0; i < b.N; i++ {
		fmt.Sprint(data)
	}
}

func BenchmarkJsonMarshal2(b *testing.B) {
	data := MakeDataDict2()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

func BenchmarkTypesDictString2(b *testing.B) {
	data := MakeDataDict2()
	for i := 0; i < b.N; i++ {
		types.String(data)
	}
}

func BenchmarkTypesMapString2(b *testing.B) {
	data := types.MapReflect(MakeDataDict2())
	for i := 0; i < b.N; i++ {
		types.String(data)
	}
}

func BenchmarkFmtSprint3(b *testing.B) {
	data := MakeDataDict3()
	for i := 0; i < b.N; i++ {
		fmt.Sprint(data)
	}
}

func BenchmarkJsonMarshal3(b *testing.B) {
	data := MakeDataDict3()
	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(data)
	}
}

func BenchmarkTypesDictString3(b *testing.B) {
	data := MakeDataDict3()
	for i := 0; i < b.N; i++ {
		types.String(data)
	}
}

func BenchmarkTypesMapString3(b *testing.B) {
	data := types.MapReflect(MakeDataDict3())
	for i := 0; i < b.N; i++ {
		types.String(data)
	}
}
