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

func BenchmarkTypesStringBytes(b *testing.B) {
	x := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}
	for i := 0; i < b.N; i++ {
		_ = types.StringBytes(x)
	}
}

func BenchmarkFmtSprintSlice(b *testing.B) {
	data := MakeDataSlice()
	for i := 0; i < b.N; i++ {
		fmt.Sprint(data)
	}
}

func BenchmarkJsonMarshalSlice(b *testing.B) {
	data := MakeDataSlice()
	for i := 0; i < b.N; i++ {
		json.Marshal(data)
	}
}

func BenchmarkTypesStringSliceGeneric(b *testing.B) {
	data := MakeDataSlice()
	for i := 0; i < b.N; i++ {
		types.StringI(data)
	}
}

func BenchmarkTypesStringSliceOptimal(b *testing.B) {
	data := MakeDataSlice()
	for i := 0; i < b.N; i++ {
		types.StringSlice(data)
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

func BenchmarkTypesStringDict1(b *testing.B) {
	data := MakeDataDict1()
	for i := 0; i < b.N; i++ {
		types.StringDict(data)
	}
}

func BenchmarkTypesStringMap1(b *testing.B) {
	data := types.MapReflect(MakeDataDict1())
	for i := 0; i < b.N; i++ {
		types.StringMap(data)
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

func BenchmarkTypesStringDict2(b *testing.B) {
	data := MakeDataDict2()
	for i := 0; i < b.N; i++ {
		types.StringDict(data)
	}
}

func BenchmarkTypesStringMap2(b *testing.B) {
	data := types.MapReflect(MakeDataDict2())
	for i := 0; i < b.N; i++ {
		types.StringMap(data)
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

func BenchmarkTypesStringDict3(b *testing.B) {
	data := MakeDataDict3()
	for i := 0; i < b.N; i++ {
		types.StringDict(data)
	}
}

func BenchmarkTypesStringMap3(b *testing.B) {
	data := types.MapReflect(MakeDataDict3())
	for i := 0; i < b.N; i++ {
		types.StringMap(data)
	}
}
