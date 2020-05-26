# `taylz.io/types`
Package `types` provides a small shim layer for go standard library type system

Types aliases for `io`, `sync`, `time`, `unsafe`, and many more

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
|`fmt.Sprint`	|6446 ns/op	|384 B/op	|
|`encoding/json.Marshal`	|4843 ns/op	|320 B/op	|
|`types.String`	|3029 ns/op	|144 B/op	|

|Map Encoding Small	|Speed	|Memory	|
|---	|---	|---	|
|`fmt.Sprint`	|33095 ns/op	|1920 B/op	|
|`encoding/json.Marshal`	|29010 ns/op	|2336 B/op	|
|`types/String(Dict)`	|07260 ns/op	|0384 B/op	|
|`types/String(Map)`	|09132 ns/op	|0384 B/op	|

|Map Encoding Large	|Speed	|Memory	|
|---	|---	|---	|
|`fmt.Sprint`	|1063150 ns/op	|66238 B/op	|
|`encoding/json.Marshal`	|0950726 ns/op	|77170 B/op	|
|`types/String(Dict)`	|0289326 ns/op	|15163 B/op	|
|`types/String(Map)`	|0291579 ns/op	|15163 B/op	|

_1. for `fmt.Sprint`, default golang encoding format is accepted, all others use JSON encoding format_

_2. for more information, see [BENCHMARKS.md](BENCHMARKS.md)_
