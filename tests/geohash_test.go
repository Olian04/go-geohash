package tests

import (
	"testing"

	"github.com/Olian04/go-geohash/geohash"
)

func TestFromString(t *testing.T) {
	hash := "ekekekekekek"
	geohash, err := geohash.FromString(hash)
	if err != nil {
		t.Fatal(err)
	}
	if geohash.ToString() != hash {
		t.Errorf("Expected hash: %s, but got hash: %s", hash, geohash.ToString())
	}
}

func TestReduceAccuracy(t *testing.T) {
	hash := "ekekekekekek"
	geohash, err := geohash.FromString(hash)
	if err != nil {
		t.Fatal(err)
	}
	reduced, err := geohash.ReduceAccuracy(2)
	if err != nil {
		t.Fatal(err)
	}
	expected := "ekekekekek"
	if reduced.ToString() != expected {
		t.Errorf("Expected hash: %s, but got hash: %s", expected, reduced.ToString())
	}
}

func TestExpandAccuracy(t *testing.T) {
	hash := "ekekekekeke"
	geohash, err := geohash.FromString(hash)
	if err != nil {
		t.Fatal(err)
	}
	expanded, err := geohash.ExpandAccuracy("k")
	if err != nil {
		t.Fatal(err)
	}
	expected := "ekekekekekek"
	if expanded.ToString() != expected {
		t.Errorf("Expected hash: %s, but got hash: %s", expected, expanded.ToString())
	}
}
