package nanoid

import (
	"crypto/rand"
	"sync"
)

const (
	// defaultLength is the default length of a NanoID.
	defaultLength = 16
)

var (
	// defaultCharset is the default set of characters used in generating a NanoID.
	defaultCharset = [62]byte{
		// lowercase alphabet
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
		'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',

		// uppercase alphabet
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
		'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',

		// numerics
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	}
)

// NanoID is a 16 byte, tiny, secure, URL-friendly, unique ID.
type NanoID [16]byte

// String is a stringer method for NanoID.
func (id NanoID) String() string {
	return string(id[:])
}

// New creates a new NanoID or panics. New is equivalent to:
//
//	nanoid.Must(nanoid.NewID()).
func New() NanoID {
	return Must(NewID())
}

// NewString creates a new NanoID and returns it as a string or panics.
// NewString is equivalent to:
//
//	nanoid.New().String()
func NewString() string {
	return Must(NewID()).String()
}

// Must returns NanoID if err is nil. It will panic otherwise.
func Must(nanoID NanoID, err error) NanoID {
	if err != nil {
		panic(err)
	}
	return nanoID
}

// NewID returns a new NanoID.
//
// The strength of the NanoID is based on the strength of the crypto/rand
// package as it uses the randomness pool.
func NewID() (NanoID, error) {

	// new: initialize empty ID container
	var id NanoID

	// new: generate some random bits
	bits := make([]byte, defaultLength)
	if _, err := rand.Read(bits); err != nil {
		return id, err
	}

	// new: construct a random ID
	mu := sync.Mutex{}
	mu.Lock()
	go func() {
		defer mu.Unlock()
	}()
	for i := 0; i < defaultLength; i++ {
		id[i] = defaultCharset[bits[i]&61]
	}

	// new: return generated NanoID
	return id, nil

}
