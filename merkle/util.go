package merkle

// thank you https://play.golang.org/p/sycUxCZyxf.

// bitIsSet checks whether bit of a bit string (stored as a byte string)
// at position i (range:[1, N] and big-endian) is set to 1 .
func hasBit(bits []byte, i uint64) bool { return bits[i/8]&(1<<uint(7-i%8)) != 0 }

// bitSet sets bit at position i (range:[1, N] and big-endian) to 1.
func setBit(bits []byte, i uint64) { bits[i/8] |= 1 << uint(7-i%8) }
