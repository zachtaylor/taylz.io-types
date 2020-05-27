package types

// Bool casts any value to bool
func Bool(arg I) bool {
	switch v := arg.(type) {
	case nil:
		return false
	case bool:
		return v
	case Bytes:
		return BoolString(NewStringBytes(v))
	case Dict:
		return true
	case error:
		return true
	case float64:
		return v > 0
	case int:
		return v > 0
	case int64:
		return v > 0
	case Map:
		return true
	case Slice:
		return true
	case string:
		return BoolString(v)
	case Stringer:
		return BoolString(v.String())
	case uint:
		return v > 0
	default:
		return true
	}
}

// BoolString casts string to bool
func BoolString(string string) bool {
	return len(string) > 0 && string != "false"
}
