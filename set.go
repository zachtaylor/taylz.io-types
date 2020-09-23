package types

// Set is a threadsafe map any->bool
type Set struct {
	cache map[I]bool
	lock  Mutex
}

// NewSet returns an empty Set
func NewSet() *Set {
	return &Set{
		cache: make(map[I]bool),
	}
}

// Has returns whether the set has the key present
func (set *Set) Has(i I) bool { return set.cache[i] }

// Add saves an element to this Set
func (set *Set) Add(i I) {
	set.lock.Lock()
	set.cache[i] = true
	set.lock.Unlock()
}

// Remove deletes an element from this Set
func (set *Set) Remove(i I) {
	set.lock.Lock()
	delete(set.cache, i)
	set.lock.Unlock()
}

// SetString is a threadsafe map string->bool
type SetString struct {
	cache map[string]bool
	lock  Mutex
}

// NewSetString returns an empty SetString
func NewSetString() *SetString {
	return &SetString{
		cache: make(map[string]bool),
	}
}

// Add saves an element to this Set
func (set *SetString) Add(string string) {
	set.lock.Lock()
	set.cache[string] = true
	set.lock.Unlock()
}

// Has returns whether the set has the string present
func (set *SetString) Has(string string) bool { return set.cache[string] }

// Slice returns a new SliceString with all elements randomly ordered
func (set *SetString) Slice() SliceString {
	set.lock.Lock()
	keys := make(SliceString, len(set.cache))
	i := 0
	for k := range set.cache {
		keys[i] = k
	}
	set.lock.Unlock()
	return keys
}

// Remove deletes an element from this Set
func (set *SetString) Remove(string string) {
	set.lock.Lock()
	delete(set.cache, string)
	set.lock.Unlock()
}

// SetUI is a threadsafe map uint->bool
type SetUI struct {
	cache map[uint]bool
	lock  Mutex
}

// NewSetUI returns an empty SetUI
func NewSetUI() *SetUI {
	return &SetUI{
		cache: make(map[uint]bool),
	}
}

// Add saves an uint to this Set
func (set *SetUI) Add(uint uint) {
	set.lock.Lock()
	set.cache[uint] = true
	set.lock.Unlock()
}

// Has returns whether the set has the uint present
func (set *SetUI) Has(uint uint) bool { return set.cache[uint] }

// Slice returns a new SliceUI with all elements randomly ordered
func (set *SetUI) Slice() SliceUI {
	set.lock.Lock()
	keys := make(SliceUI, len(set.cache))
	i := 0
	for k := range set.cache {
		keys[i] = k
	}
	set.lock.Unlock()
	return keys
}

// Remove deletes an uint from this Set
func (set *SetUI) Remove(uint uint) {
	set.lock.Lock()
	delete(set.cache, uint)
	set.lock.Unlock()
}
