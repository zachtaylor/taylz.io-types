package types

import (
	"fmt"
	"strconv"
	"strings"
)

// String casts any value to string
func String(arg I) string {
	switch v := arg.(type) {
	case nil:
		return ""
	case bool:
		return StringBool(v)
	case Bytes:
		return StringBytes(v)
	case Dict:
		return StringDict(v)
	case error:
		return `error: "` + v.Error() + `"`
	case float64:
		return StringF64(v)
	case int:
		return StringInt(v)
	case int64:
		return StringI64(v)
	case Map:
		return StringMap(v)
	case Slice:
		return StringSlice(v)
	case string:
		return v
	case Stringer:
		return v.String()
	case uint:
		return StringUint(v)
	default:
		return fmt.Sprint(arg)
	}
}

// StringBool casts bool to string
func StringBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

// StringBytes casts []byte to string
func StringBytes(b []byte) string { return *(*string)(Pointer(&b)) }

// StringContains uses strings.Contains
func StringContains(a, b string) bool { return strings.Contains(a, b) }

// StringDict casts Dict to string
func StringDict(dict Dict) (str string) {
	sb := poolStringBuilder.Get().(*StringBuilder)
	sb.Grow(ceil32(ByteSizeJSONDict(dict)))
	EncodeJSONDict(sb, dict)
	str = sb.String()
	sb.Reset()
	poolStringBuilder.Put(sb)
	return
}

// StringF64 casts float64 to string
func StringF64(f float64) string { return strconv.FormatFloat(f, 'f', -1, 64) }

// StringInCharset returns len(Trim(a, b)) < 1
func StringInCharset(a, b string) bool { return len(StringTrim(a, b)) < 1 }

// StringInt casts int to string
func StringInt(i int) string { return strconv.Itoa(i) }

// StringI64 casts int64 to string
func StringI64(i64 int64) string { return strconv.FormatInt(i64, 10) }

// StringJoin is an alias for strings.Join
func StringJoin(a []string, sep string) string { return strings.Join(a, sep) }

// StringLastIndex is an alias for strings.LastIndex
func StringLastIndex(a, b string) int { return strings.LastIndex(a, b) }

// StringMap casts Map to string
func StringMap(m Map) (str string) {
	sb := poolStringBuilder.Get().(*StringBuilder)
	sb.Grow(ceil32(ByteSizeJSONMap(m)))
	EncodeJSONMap(sb, m)
	str = sb.String()
	sb.Reset()
	poolStringBuilder.Put(sb)
	return
}

// StringOutCharset returns len(Trim()) > 0
func StringOutCharset(a, b string) bool { return len(StringTrim(a, b)) > 0 }

// StringSlice casts Slice to string
func StringSlice(slice Slice) (str string) {
	sb := poolStringBuilder.Get().(*StringBuilder)
	sb.Grow(ceil32(ByteSizeJSONSlice(slice)))
	EncodeJSONSlice(sb, slice)
	str = sb.String()
	sb.Reset()
	poolStringBuilder.Put(sb)
	return
}

// StringSplit is an alias for strings.Split
func StringSplit(s, sep string) []string { return strings.Split(s, sep) }

// StringTrim uses strings.Trim
func StringTrim(a, b string) string { return strings.Trim(a, b) }

// StringUint casts uint to string
func StringUint(ui uint) string { return StringInt(int(ui)) }
