package merkle

import (
	"math/big"
)

// bitIsSet checks whether bit of a bit string (stored as a byte string)
// at position i (range:[1, N] and big-endian) is set to 1 .
func bitIsSet(bits []byte, i uint64) bool { return bits[i/8]&(1<<uint(7-i%8)) != 0 }

// bitSet sets bit at position i (range:[1, N] and big-endian) to 1.
func bitSet(bits []byte, i uint64) { bits[i/8] |= 1 << uint(7-i%8) }

func bigAddInt(x *big.Int, y int64) *big.Int    { return new(big.Int).Add(x, big.NewInt(y)) }
func bigAddBig(x *big.Int, y *big.Int) *big.Int { return new(big.Int).Add(x, y) }
func bigSubInt(x *big.Int, y int64) *big.Int    { return new(big.Int).Sub(x, big.NewInt(y)) }
func bigSubBig(x *big.Int, y *big.Int) *big.Int { return new(big.Int).Sub(x, y) }
func bigDivInt(x *big.Int, y int64) *big.Int    { return new(big.Int).Div(x, big.NewInt(y)) }
func bigDivBig(x *big.Int, y *big.Int) *big.Int { return new(big.Int).Div(x, y) }
func bigModInt(x *big.Int, y int64) *big.Int    { return new(big.Int).Mod(x, big.NewInt(y)) }
func bigModBig(x *big.Int, y *big.Int) *big.Int { return new(big.Int).Mod(x, y) }
