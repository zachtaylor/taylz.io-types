package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

// JSON returns JSON string encoding of an arg
func JSON(arg I) (str string) {
	sb := poolStringBuilder.Get().(*StringBuilder)
	sb.Grow(ceil32(ByteSizeJSON(arg)))
	EncodeJSON(sb, arg)
	str = sb.String()
	sb.Reset()
	poolStringBuilder.Put(sb)
	return
}

// EncodeJSON writes arg to *StringBuilder in JSON
func EncodeJSON(sb *StringBuilder, arg I) {
	const charquote = '"'
	switch v := arg.(type) {
	case nil:
		sb.WriteString("null")
	case bool:
		sb.WriteString(NewStringBool(v))
	case Bytes:
		sb.WriteByte(charquote)
		sb.WriteString(jsonEscapeString.Replace(NewStringBytes(v)))
		sb.WriteByte(charquote)
	case Dict:
		EncodeJSONDict(sb, v)
	case error:
		sb.WriteString(`"error: ` + jsonEscapeString.Replace(v.Error()) + `"`)
	case float64:
		sb.WriteString(NewStringF64(v))
	case int:
		sb.WriteString(NewStringInt(v))
	case int64:
		sb.WriteString(NewStringI64(v))
	case Map:
		EncodeJSONMap(sb, v)
	case Slice:
		EncodeJSONSlice(sb, v)
	case string:
		sb.WriteByte(charquote)
		sb.WriteString(jsonEscapeString.Replace(v))
		sb.WriteByte(charquote)
	case Stringer:
		sb.WriteString(v.String())
	default:
		val := Reflect(arg)
		k := val.Kind()
		if k == SliceKind {
			EncodeJSONSlice(sb, sliceValue(val))
		} else if k == MapKind {
			EncodeJSONMap(sb, mapValue(val))
		} else {
			fmt.Fprint(sb, arg)
		}
	}
}

// EncodeJSONDict writes Dict to *StringBuilder in JSON
func EncodeJSONDict(sb *StringBuilder, dict Dict) {
	const charopen = '{'
	const chardef = ':'
	const charquote = '"'
	const charsep = ','
	const charclose = '}'
	sb.WriteByte(charopen)
	first := true
	for k, v := range dict { // k type not required, but golang requires this distinction
		if !first {
			sb.WriteByte(charsep)
		} else {
			first = false
		}
		sb.WriteByte(charquote)
		sb.WriteString(jsonEscapeString.Replace(k))
		sb.WriteByte(charquote)
		sb.WriteByte(chardef)
		EncodeJSON(sb, v)
	}
	sb.WriteByte(charclose)
}

// EncodeJSONMap writes Map to *StringBuilder in JSON
func EncodeJSONMap(sb *StringBuilder, m Map) {
	const charopen = '{'
	const chardef = ':'
	const charsep = ','
	const charclose = '}'
	sb.WriteByte(charopen)
	first := true
	for k, v := range m { // k type not required, but golang requires this distinction
		if !first {
			sb.WriteByte(charsep)
		} else {
			first = false
		}
		EncodeJSON(sb, k)
		sb.WriteByte(chardef)
		EncodeJSON(sb, v)
	}
	sb.WriteByte(charclose)
}

// EncodeJSONSlice writes Slice to *StringBuilder in JSON
func EncodeJSONSlice(sb *StringBuilder, slice Slice) {
	const charopen = '['
	const charsep = ','
	const charclose = ']'
	sb.WriteByte(charopen)
	for k, v := range slice {
		if k > 0 {
			sb.WriteByte(charsep)
		}
		EncodeJSON(sb, v)
	}
	sb.WriteByte(charclose)
}

// jsonEscapeStringReplacer is used by jsonEscapeString
var jsonEscapeString = strings.NewReplacer(`\`, `\\`, `"`, `\"`)

// DecodeJSON returns json.NewDecoder().Decode()
func DecodeJSON(r Reader, v I) error { return json.NewDecoder(r).Decode(v) }
