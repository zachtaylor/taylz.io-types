package types

import "reflect"

// Kind = reflect.Kind
type Kind = reflect.Kind

// Value = reflect.Value
type Value = reflect.Value

// SliceKind = reflect.Slice
const SliceKind = reflect.Slice

// MapKind = reflect.Map
const MapKind = reflect.Map

// T = reflect.Type
type T = reflect.Type

// Get returns reflect.TypeOf
func Get(arg I) T { return reflect.TypeOf(arg) }

// KindOf is an alias for Get().Kind()
func KindOf(arg I) Kind { return Get(arg).Kind() }

// Reflect returns reflect.ValueOf
func Reflect(arg I) Value { return reflect.ValueOf(arg) }
