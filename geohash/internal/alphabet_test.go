package internal

import (
	"testing"
)

func TestDecode(t *testing.T) {
	for index, c := range Base32ghs {
		i, err := Decode(c)
		if err != nil {
			t.Fatal(err)
		}
		if i != index {
			t.Fatalf("expected %d, got %d", index, i)
		}
	}
}

func TestEncode(t *testing.T) {
	for index, c := range Base32ghs {
		r, err := Encode(index)
		if err != nil {
			t.Fatal(err)
		}
		if r != c {
			t.Fatalf("expected %c, got %c", c, r)
		}
	}
}
