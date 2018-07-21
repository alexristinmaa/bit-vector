package bitvector

import "testing"

// TestAnd tests the and function
func TestAnd(t *testing.T) {
  a := new(BitVector).Add(1,3,5,7)
  b := new(BitVector).Add(0,2,4,7)
  a.And(b)
  if a.AsBinary() != "0000000000000000000000000000000000000000000000000000000010000000" {
    t.Error("Expected 0000000000000000000000000000000000000000000000000000000010000000, got", a.AsBinary())
  }
}

// TestOr tests the or function
func TestOr(t *testing.T) {
  a := new(BitVector).Add(1,3,5,7)
  b := new(BitVector).Add(0,2,4,6)
  a.Or(b)
  if a.AsBinary() != "0000000000000000000000000000000000000000000000000000000011111111" {
    t.Error("Expected 0000000000000000000000000000000000000000000000000000000011111111, got", a.AsBinary())
  }
}

// TestXor tests the xor function
func TestXor(t *testing.T) {
  a := new(BitVector).Add(0,2,3,4)
  b := new(BitVector).Add(0,1,2,3)

  a.Xor(b)

  if a.AsBinary() != "0000000000000000000000000000000000000000000000000000000000010010" {
    t.Error("Expected 0000000000000000000000000000000000000000000000000000000000010010, got", a.AsBinary())
  }
}

// TestClear tests the clear function
func TestClear(t *testing.T) {
  a := new(BitVector).Add(0,1,2,3,4,5,6,7,8,9)

  a.Clear()

  if *a != 0 {
    t.Error("Expected 0000000000000000000000000000000000000000000000000000000000010010, got", a.AsBinary())
  }
}

// TestShiftL tests the ShiftL function
func TestShiftL(t *testing.T) {
  a := new(BitVector).Add(0,1,2,3,4,5,6,7,8,9)

  a.ShiftL(2)

  if a.AsBinary() != "0000000000000000000000000000000000000000000000000000111111111100" {
    t.Error("Expected 0000000000000000000000000000000000000000000000000000111111111100, got", a.AsBinary())
  }
}

// TestShiftR tests the ShiftR function
func TestShiftR(t *testing.T) {
  a := new(BitVector).Add(0,1,2,3,4,5,6,7,8,9)

  a.ShiftR(2)

  if a.AsBinary() != "0000000000000000000000000000000000000000000000000000000011111111" {
    t.Error("Expected 0000000000000000000000000000000000000000000000000000000011111111, got", a.AsBinary())
  }
}
