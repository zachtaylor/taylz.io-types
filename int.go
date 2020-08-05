package types

import "strconv"

// IntString casts string to int
func IntString(s string) int {
	i, _ := strconv.ParseInt(s, 10, 0)
	return int(i)
}

// Int casts any type to int
func Int(i I) int {
	switch v := i.(type) {
	case bool:
		if v {
			return 1
		} else {
			return -1
		}
	case int:
		return v
	case uint:
		return int(v)
	case int64:
		return int(v)
	case uint64:
		return int(v)
	case float64:
		return int(v)
	case string:
		return IntString(v)
	case Stringer:
		return IntString(v.String())
	default:
		return 0
	}
}
