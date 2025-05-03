package geohash

import (
	"testing"
)

func TestFromString(t *testing.T) {
	hash := "ekekekekekek"
	geohash, err := FromString(hash)
	if err != nil {
		t.Fatal(err)
	}
	if geohash.ToString() != hash {
		t.Errorf("Expected hash: %s, but got hash: %s", hash, geohash.ToString())
	}
}

func TestReduceAccuracy(t *testing.T) {
	hash := "ekekekekekek"
	geohash, err := FromString(hash)
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

func TestIncreaseAccuracy(t *testing.T) {
	hash := "ekekekekeke"
	geohash, err := FromString(hash)
	if err != nil {
		t.Fatal(err)
	}
	expanded, err := geohash.IncreaseAccuracy("k")
	if err != nil {
		t.Fatal(err)
	}
	expected := "ekekekekekek"
	if expanded.ToString() != expected {
		t.Errorf("Expected hash: %s, but got hash: %s", expected, expanded.ToString())
	}
}
