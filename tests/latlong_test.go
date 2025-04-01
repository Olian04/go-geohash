package tests

import (
	"testing"

	"github.com/Olian04/go-geohash/geohash/lib"
	"gonum.org/v1/gonum/floats"
)

func TestToLatLong(t *testing.T) {
	lat, long, err := lib.ToLatLong("ekekekekekek")
	if err != nil {
		t.Fatal(err)
	}
	latLong := []float64{lat, long}
	expectedLatLong := []float64{26.12903223, -29.03225804}
	if !floats.EqualApprox(latLong, expectedLatLong, 1e-8) {
		t.Errorf("Expected lat: %f, long: %f, but got lat: %f, long: %f", expectedLatLong[0], expectedLatLong[1], latLong[0], latLong[1])
	}
}

func TestFromLatLong(t *testing.T) {
	hash := lib.FromLatLong(26.12903223, -29.03225804)
	expectedHash := "ekekekekekek"
	if hash != expectedHash {
		t.Errorf("Expected hash: %s, but got hash: %s", expectedHash, hash)
	}
}
