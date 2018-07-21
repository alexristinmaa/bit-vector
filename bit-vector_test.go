package bitvector

import "testing"

// TestAdd tests the add function
func TestAdd(t *testing.T) {
  b := new(BitVector)
  b.Add(2)
  if !b.Has(2) {
    t.Error("Expected true, got false")
  }
  b.Add(0, 4, 5)
  if !b.Has(0, 4, 5) {
    t.Error("Expected true, got false")
  }
  if b.Has(1,2,3) {
    t.Error("Expected false, got true")
  }
}

// TestRemove tests the remove function
func TestRemove(t *testing.T) {
  b := new(BitVector)

  b.Add(3).Add(4).Add(10).Add(30).Remove(4)

  if b.Has(4) {
    t.Error("Expected false, got true")
  }

  b.Remove(3,10)

  if b.Has(3,10) {
    t.Error("Expected false, got true")
  }

}

// TestHas tests the has function
func TestHas(t *testing.T) {
  b := new(BitVector)

  b.Add(4).Add(23).Add(3).Add(1)
  v := b.Has(23)

  if !v {
    t.Error("Expected true, got false")
  }

  v = b.Has(30)

  if v {
    t.Error("Expected false, got true")
  }
}

// TestAsBinary tests the AsBinary function
func TestAsBinary(t *testing.T) {
  b := new(BitVector)

  b.Add(0).Add(1).Add(2).Add(3).Add(4).Add(10)

  v := b.AsBinary()

  if v != "0000000000000000000000000000000000000000000000000000010000011111" {
    t.Error("Expected 0000000000000000000000000000000000000000000000000000010000011111, got", v)
  }

  b.Remove(10)

  v = b.AsBinary()

  if v != "0000000000000000000000000000000000000000000000000000000000011111" {
    t.Error("Expected 0000000000000000000000000000000000000000000000000000000000011111, got", v)
  }
}

// TestEquals tests the equals function
func TestEquals(t *testing.T) {
  a := new(BitVector).Add(1,2,3,4,5,6,7,8,9,30)
  b := new(BitVector).Add(1,2,3,4,5,6,7,8,9,30)

  if !a.Equals(b) {
    t.Error("Expected true, got false")
  }

  b.Remove(30)

  if a.Equals(b) {
    t.Error("Expected false, got true")
  }
}
