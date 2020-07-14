package types

import "strconv"

// IntString casts string to int
func IntString(s string) int {
	i, _ := strconv.ParseInt(s, 10, 0)
	return int(i)
}
