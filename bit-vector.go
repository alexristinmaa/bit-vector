// Bit-vector is a bit vector implementation in go
// It is a simple implementation for anyone who wants to use bit vectors/arrays with a basic range of usage 
// but still requires more specific operations such as bitwise and equals.
package bitvector

// BitVector is a bit array
type BitVector uint64

// Add adds multiple values to the bit array
func (b *BitVector) Add(nums ...uint) *BitVector {
  for _, n := range nums {
    b.add(n)
  }
  return b
}

// Has checks if the bit vector has an unsigned integer in it
func (b *BitVector) Has(nums ...uint) bool {
	for _, n := range nums {
    if(!b.has(n)) {
      return false
    }
  }

  return true
}

// Remove removes multiple values from the bit vector
func (b *BitVector) Remove(nums ...uint) *BitVector {
  for _, n := range nums {
    b.remove(n)
  }

  return b
}

// AsBinary returns the vector as a 64 character string in which it is
// represented in binary
func (b *BitVector) AsBinary() string {
  res := ""
  n := *b
  for n > 0 {
    if n % 2 == 0 {
      res = "0" + res
    } else {
      res = "1" + res
    }
    n /= 2
  }

  leading := ""

  for i := 0; i < 64-len(res); i++ {
    leading += "0"
  }

  return leading + res
}

// Equals checks if to bit vectors are equal
func (b *BitVector) Equals(other *BitVector) bool {
  return other.AsBinary() == b.AsBinary()
}

// And ands (&) two bit vectors
func (b *BitVector) And(other *BitVector) *BitVector {
  *b &= *other

  return b
}

// Or ors (|) two bit vectors
func (b *BitVector) Or(other *BitVector) *BitVector {
  *b |= *other

  return b
}

// Xor xors (^) the two bit vectors
func (b *BitVector) Xor(other *BitVector) *BitVector {
  *b ^= *other

  return b
}

// Clear clears the entire bit vector
func (b *BitVector) Clear() *BitVector {
  *b = 0

  return b
}
/*
_    _ _   _ ________   _______   ____  _____ _______ ______ _____
| |  | | \ | |  ____\ \ / /  __ \ / __ \|  __ \__   __|  ____|  __ \
| |  | |  \| | |__   \ V /| |__) | |  | | |__) | | |  | |__  | |  | |
| |  | | . ` |  __|   > < |  ___/| |  | |  _  /  | |  |  __| | |  | |
| |__| | |\  | |____ / . \| |    | |__| | | \ \  | |  | |____| |__| |
\____/|_| \_|______/_/ \_\_|     \____/|_|  \_\ |_|  |______|_____/
*/

// has checks if a number exists in the bit BitVector
func (b *BitVector) has(n uint) bool {
  return *b & BitVector(uint64(1) << n) > 0
}

// remove removes a value from the bit vector
func (b *BitVector) remove(n uint) *BitVector {
  *b &^= BitVector(uint64(1) << n)
  return b
}

// add adds an unsigned integer to the bit array
func (b *BitVector) add(n uint) *BitVector {
  *b |= BitVector((uint64(1) << n))
  return b
}
