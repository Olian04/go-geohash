package tests

import (
	"testing"

	"github.com/Olian04/go-geohash/geohash/lib"
)

func TestDecode(t *testing.T) {
	for index, c := range lib.Base32ghs {
		i, err := lib.Decode(c)
		if err != nil {
			t.Fatal(err)
		}
		if i != index {
			t.Fatalf("expected %d, got %d", index, i)
		}
	}
}

func TestEncode(t *testing.T) {
	for index, c := range lib.Base32ghs {
		r, err := lib.Encode(index)
		if err != nil {
			t.Fatal(err)
		}
		if r != c {
			t.Fatalf("expected %c, got %c", c, r)
		}
	}
}
