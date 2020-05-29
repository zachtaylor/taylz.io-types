package types

// Map is a map of I -> I
type Map = map[I]I

// MapReflect returns a Map from an arg
func MapReflect(arg I) (out Map) {
	if val := Reflect(arg); val.Kind() == KindMap {
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
