package geohash

import (
	"fmt"

	"github.com/Olian04/go-geohash/geohash/internal"
)

type GeoHash struct {
	hash string
}

// FromString creates a new GeoHash from a string
func FromString(s string) (GeoHash, error) {
	if !internal.ValidateLength(s) {
		return GeoHash{}, fmt.Errorf("geohash: %s is not of valid length", s)
	}
	if !internal.ValidateAlphabet(s) {
		return GeoHash{}, fmt.Errorf("geohash: %s contains invalid characters", s)
	}
	return GeoHash{hash: s}, nil
}

// FromLatLong creates a new GeoHash from a latitude and longitude
func FromLatLong(lat float64, long float64) (GeoHash, error) {
	hash := internal.FromLatLong(lat, long)
	return FromString(hash)
}

// EqualizeAccuracy equalizes the accuracy of two geohashes
// It returns the two geohashes with the same accuracy by truncating the longer of the two hashes
func EqualizeAccuracy(hash1 string, hash2 string) (GeoHash, GeoHash, error) {
	targetLen := min(len(hash2), len(hash1))
	h1, err := FromString(hash1[:targetLen])
	if err != nil {
		return GeoHash{}, GeoHash{}, err
	}
	h2, err := FromString(hash2[:targetLen])
	if err != nil {
		return GeoHash{}, GeoHash{}, err
	}
	return h1, h2, nil
}

// CapAccuracy caps the accuracy of the geohash to the given level
func (g GeoHash) CapAccuracy(accuracyLevel int) GeoHash {
	if accuracyLevel < 0 {
		return g
	}
	if accuracyLevel > len(g.hash) {
		return g
	}
	hash := g.hash[:accuracyLevel]
	geoHash, err := FromString(hash)
	if err != nil {
		panic(err)
	}
	return geoHash
}

// ReduceAccuracy reduces the accuracy of the geohash by the given number of levels
func (g GeoHash) ReduceAccuracy(levels int) (GeoHash, error) {
	if levels < 0 {
		return GeoHash{}, fmt.Errorf("geohash: accuracy levels cannot be negative")
	}
	if levels > len(g.hash) {
		return GeoHash{}, fmt.Errorf("geohash: accuracy levels cannot be greater than the length of the hash")
	}
	hash := g.hash[:len(g.hash)-levels]
	return FromString(hash)
}

// IncreaseAccuracy increases the accuracy of the geohash by the given characters
func (g GeoHash) IncreaseAccuracy(characters string) (GeoHash, error) {
	if len(characters) == 0 {
		return GeoHash{}, fmt.Errorf("geohash: characters cannot be empty")
	}
	hash := g.hash + characters
	return FromString(hash)
}

// Accuracy returns the accuracy of the geohash in meters
func (g GeoHash) Accuracy() float64 {
	return internal.Accuracy(g.hash)
}

// AccuracyLevel returns the accuracy level of the geohash
func (g GeoHash) AccuracyLevel() int {
	return len(g.hash)
}

// ApproximateDistanceTo returns the approximate distance between two geohashes in meters
// The distance is approximate because it is based on the index of the first
// non-matching character. Table of distance per character index:
//
//	0: 20_000_000
//	1: 5_003_530
//	2: 625_441
//	3: 123_264
//	4: 19_545
//	5: 3_803
//	6: 610
//	7: 118
//	8: 19
//	9: 3.71
//	10: 0.6
//	11: 0.09
//	12: 0.02
func (g GeoHash) ApproximateDistanceTo(other GeoHash) float64 {
	return internal.ApproximateDistance(g.hash, other.hash)
}

// ToString returns the string representation of the geohash
func (g GeoHash) ToString() string {
	return g.hash
}

// ToLatLong returns the latitude and longitude of the geohash
func (g GeoHash) ToLatLong() (float64, float64, error) {
	lat, long, err := internal.ToLatLong(g.hash)
	if err != nil {
		return 0, 0, err
	}
	return lat, long, nil
}
