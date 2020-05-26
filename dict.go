package types

// Dict is a map with string key
type Dict = map[string]interface{}

// GetDictKeys returns an alphabetically sorted slice of keys from a Dict
func GetDictKeys(dict Dict) []string {
	keys := make(SliceString, len(dict))
	var i int
	for k := range dict {
		keys[i] = k
		i++
	}
	keys.Sort()
	return keys
}
