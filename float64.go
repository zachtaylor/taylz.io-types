package types

import "strconv"

// F64String casts string to float64
func F64String(s string) (f float64) { f, _ = strconv.ParseFloat(s, 64); return }
