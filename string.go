package types

import (
	"fmt"
	"strconv"
	"strings"
)

// NewString casts any value to string
func NewString(arg I) string {
	switch v := arg.(type) {
	case nil:
		return ""
	case bool:
		return NewStringBool(v)
	case Bytes:
		return NewStringBytes(v)
	case Dict:
		return NewStringDict(v)
	case error:
		return `error: "` + v.Error() + `"`
	case float64:
		return NewStringF64(v)
	case int:
		return NewStringInt(v)
	case int64:
		return NewStringI64(v)
	case Map:
		return NewStringMap(v)
	case Slice:
		return NewStringSlice(v)
	case string:
		return v
	case Stringer:
		return v.String()
	case uint:
		return NewStringUint(v)
	default:
		return fmt.Sprint(arg)
	}
}

// NewStringBool casts bool to string
func NewStringBool(b bool) string {
	const true = "true"
	const false = "false"
	if b {
		return true
	}
	return false
}

// NewStringBytes casts []byte to string
func NewStringBytes(b []byte) string { return *(*string)(Pointer(&b)) }

// NewStringDict casts Dict to string
func NewStringDict(dict Dict) (str string) {
	sb := poolStringBuilder.Get().(*StringBuilder)
	sb.Grow(ceil32(ByteSizeJSONDict(dict)))
	EncodeJSONDict(sb, dict)
	str = sb.String()
	sb.Reset()
	poolStringBuilder.Put(sb)
	return
}

// NewStringF64 casts float64 to string
func NewStringF64(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) }

// NewStringInt casts int to string
func NewStringInt(i int) string { return strconv.Itoa(i) }

// NewStringI64 casts int64 to string
func NewStringI64(i64 int64) string { return strconv.FormatInt(i64, 10) }

// NewStringJoin is an alias for strings.Join
func NewStringJoin(a []string, sep string) string { return strings.Join(a, sep) }

// NewStringMap casts Map to string
func NewStringMap(m Map) (str string) {
	sb := poolStringBuilder.Get().(*StringBuilder)
	sb.Grow(ceil32(ByteSizeJSONMap(m)))
	EncodeJSONMap(sb, m)
	str = sb.String()
	sb.Reset()
	poolStringBuilder.Put(sb)
	return
}

// NewStringSlice casts Slice to string
func NewStringSlice(slice Slice) (str string) {
	sb := poolStringBuilder.Get().(*StringBuilder)
	sb.Grow(ceil32(ByteSizeJSONSlice(slice)))
	EncodeJSONSlice(sb, slice)
	str = sb.String()
	sb.Reset()
	poolStringBuilder.Put(sb)
	return
}

// NewStringTrim uses strings.Trim
func NewStringTrim(a, b string) string { return strings.Trim(a, b) }

// NewStringUint casts uint to string
func NewStringUint(ui uint) string { return NewStringInt(int(ui)) }

// StringContains uses strings.Contains
func StringContains(a, b string) bool { return strings.Contains(a, b) }

// StringInCharset returns len(Trim(a, b)) < 1
func StringInCharset(a, b string) bool { return len(NewStringTrim(a, b)) < 1 }

// StringLastIndex is an alias for strings.LastIndex
func StringLastIndex(a, b string) int { return strings.LastIndex(a, b) }

// StringOutCharset returns len(Trim()) > 0
func StringOutCharset(a, b string) bool { return len(NewStringTrim(a, b)) > 0 }
