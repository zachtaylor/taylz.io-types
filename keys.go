package types

// GetKeysDict returns a sorted slice of keys for the dict
func GetKeysDict(dict Dict) SliceString {
	keys := make(SliceString, len(dict))
	var i int
	for k := range dict {
		keys[i] = k
		i++
	}
	keys.Sort()
	return keys
}

// GetKeysMap returns a slice of keys for the map
func GetKeysMap(m Map) Slice {
	keys := make(Slice, len(m))
	var i int
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}
