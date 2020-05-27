# `taylz.io/types`
Package `types` provides a small shim layer for go standard library type system

Types aliases for `fmt`, `io`, `rand`, `strings`, `sync`, `time`, `unsafe`, and more

Improved JSON encoding, with average 66% speed savings, and over 75% memory savings

## New Types

### Any

```
I = interface{}
```

### Bytes

```
Bytes = []byte
```

### Dictionary

```
Dict = map[string]I
```

### Map

```
Map = map[I]I
```

### Slice

```
Slice = []I
```

### Source

```
Source struct {
	file string
	line int
}
```

## Type Conversion

Provides common type conversions, such as "to string" and "to json"

### String Conversion

String a byte slice using the `unsafe` hack for 95% speed savings

### JSON Encoding Speed

|Slice Encoding	|Speed	|Memory	|
|---	|---	|---	|
|`fmt.Sprint`	|6463 ns/op	|384 B/op	|
|`json.Marshal`	|4723 ns/op	|320 B/op	|
|`types.NewString`	|2982 ns/op	|144 B/op	|
|`types.NewStringSlice`	|2797 ns/op	|112 B/op	|

|Map Encoding Small	|Speed	|Memory	|
|---	|---	|---	|
|`fmt.Sprint`	|32371 ns/op	|1920 B/op	|
|`json.Marshal`	|28480 ns/op	|2336 B/op	|
|`types.NewStringDict`	|07159 ns/op	|0384 B/op	|
|`types.NewStringMap`	|09287 ns/op	|0384 B/op	|

|Map Encoding Large	|Speed	|Memory	|
|---	|---	|---	|
|`fmt.Sprint`	|1052620 ns/op	|66238 B/op	|
|`json.Marshal`	|0935566 ns/op	|77170 B/op	|
|`types.NewStringDict`	|0292551 ns/op	|15164 B/op	|
|`types.NewStringMap`	|0292671 ns/op	|15163 B/op	|

_1. for `fmt.Sprint`, default golang encoding format is accepted, all others use JSON encoding format_

_2. for more information, see [BENCHMARKS.md](BENCHMARKS.md)_
