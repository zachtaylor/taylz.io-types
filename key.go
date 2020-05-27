package types

// KeyStorer is an interface that tests keys
type KeyStorer interface {
	Has(string) bool
}

// KeyGen returns a new unique Key for the KeyStorer
//
// KeyStorer synchronization not within scope; call this func within applicable lock state
func KeyGen(store KeyStorer, size int, charset string, rand *Rand) string {
	bytes := make([]byte, size)
	var key string
	for done := false; !done; {
		for i, clen := 0, int32(len(charset)); i < size; i++ {
			bytes[i] = charset[int(rand.Int31n(clen))]
		}
		key = NewStringBytes(bytes)
		done = !store.Has(key)
	}
	return key
}

// KeyGenAlpha uses KeyGen with CharsetAlpha
func KeyGenAlpha(store KeyStorer, size int, rand *Rand) string {
	return KeyGen(store, size, CharsetAlpha, rand)
}

// KeyGenAlphaCapital uses KeyGen with CharsetAlphaCapital
func KeyGenAlphaCapital(store KeyStorer, size int, rand *Rand) string {
	return KeyGen(store, size, CharsetAlphaCapital, rand)
}

// KeyGenAlphaCapitalNumeric uses `KeyGen` with `CharsetAlphaCapitalNumeric`
func KeyGenAlphaCapitalNumeric(store KeyStorer, size int, rand *Rand) string {
	return KeyGen(store, size, CharsetAlphaCapitalNumeric, rand)
}

// KeyGenCapital uses `KeyGen` with `CharsetCapital`
func KeyGenCapital(store KeyStorer, size int, rand *Rand) string {
	return KeyGen(store, size, CharsetCapital, rand)
}

// KeyGenCapitalNumeric uses KeyGen with CharsetCapitalNumeric
func KeyGenCapitalNumeric(store KeyStorer, size int, rand *Rand) string {
	return KeyGen(store, size, CharsetCapitalNumeric, rand)
}

// KeyGenNumeric uses KeyGen with CharsetNumeric
func KeyGenNumeric(store KeyStorer, size int, rand *Rand) string {
	return KeyGen(store, size, CharsetNumeric, rand)
}

// KeyGenSettings stores key generation settings
type KeyGenSettings struct {
	Size    int
	Charset string
	Rand    *Rand
}

// NewKeyGenSettings returns a new KeyGenSettings
func NewKeyGenSettings(size int, charset string, rand *Rand) *KeyGenSettings {
	return &KeyGenSettings{
		Size:    size,
		Charset: charset,
		Rand:    rand,
	}
}

// KeyGen uses KeyStorer its settings to call KeyGen
//
// As with KeyGen, lock the store to guarantee maintainted uniqueness
func (keygen *KeyGenSettings) KeyGen(store KeyStorer) string {
	return KeyGen(store, keygen.Size, keygen.Charset, keygen.Rand)
}
