package types

import "math"

// ByteSizeJSON returns a guess at the number of bytes needed to encode arg in JSON
func ByteSizeJSON(arg I) int {
	switch v := arg.(type) {
	case nil:
		return 4 // len("null")
	case bool:
		if v {
			return 4 // len("true")
		}
		return 5 // len("false")
	case Bytes:
		return len(v) + 2 // add escape quotes
	case Dict:
		return ByteSizeJSONDict(v)
	case error:
		return len(v.Error()) + 2 // add escape quotes
	case float64:
		return int(math.Ceil(math.Log10(v + 1)))
	case int:
		return int(math.Ceil(math.Log10(float64(v + 1))))
	case int64:
		return int(math.Ceil(math.Log10(float64(v + 1))))
	case Map:
		return ByteSizeJSONMap(v)
	case Slice:
		return ByteSizeJSONSlice(v)
	case string:
		return len(v) + 2 // quotes
	case Stringer:
		return len(v.String())
	case uint:
		return int(math.Ceil(math.Log10(float64(v + 1))))
	default:
		return 8 // guess
	}
}

// ByteSizeJSONDict returns a guess at the number of bytes needed to encode the Dict in JSON
func ByteSizeJSONDict(dict Dict) int {
	first := true
	size := 2 // brackets
	for k, v := range dict {
		if first {
			first = false
		} else {
			size++ // comma
		}
		size += 2 + len(k) + 1 + ByteSizeJSON(v) // quotes and colon
	}
	return size
}

// ByteSizeJSONMap returns a guess at the number of bytes needed to encode the Map in JSON
func ByteSizeJSONMap(m Map) int {
	first := true
	size := 2 // brackets
	for k, v := range m {
		if first {
			first = false
		} else {
			size++ // comma
		}
		size += ByteSizeJSON(k) + 1 + ByteSizeJSON(v) // colon
	}
	return size
}

// ByteSizeJSONSlice returns a guess at the number of bytes needed to encode the Slice in JSON
func ByteSizeJSONSlice(slice Slice) int {
	first := true
	size := 2 // brackets
	for _, v := range slice {
		if first {
			first = false
		} else {
			size++ // separator
		}
		size += ByteSizeJSON(v)
	}
	return size
}
