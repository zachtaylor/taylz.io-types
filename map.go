package types

// Map is a map of I -> I
type Map = map[I]I

// MapReflect returns a Map from an arg
func MapReflect(arg I) (out Map) {
	if val := Reflect(arg); val.Kind() == MapKind {
		out = mapValue(val)
	}
	return
}

func mapValue(val Value) (out Map) {
	out = make(Map)
	for _, k := range val.MapKeys() {
		out[k.Interface()] = val.MapIndex(k).Interface()
	}
	return
}

// // MapReflect returns MapValue(Reflect())
// func MapReflect(arg I) Map { return MapValue(Reflect(arg)) }
