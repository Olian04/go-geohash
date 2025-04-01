package geohash

import (
	"fmt"

	"github.com/Olian04/go-geohash/geohash/lib"
)

type GeoHash struct {
	hash string
}

func FromString(s string) (GeoHash, error) {
	if !lib.ValidateLength(s) {
		return GeoHash{}, fmt.Errorf("geohash: %s is not of valid length", s)
	}
	if !lib.ValidateAlphabet(s) {
		return GeoHash{}, fmt.Errorf("geohash: %s contains invalid characters", s)
	}
	return GeoHash{hash: s}, nil
}

func FromLatLong(lat float64, long float64) (GeoHash, error) {
	hash := lib.FromLatLong(lat, long)
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

// ExpandAccuracy expands the accuracy of the geohash by the given characters
func (g GeoHash) ExpandAccuracy(characters string) (GeoHash, error) {
	if len(characters) == 0 {
		return GeoHash{}, fmt.Errorf("geohash: characters cannot be empty")
	}
	hash := g.hash + characters
	return FromString(hash)
}

func (g GeoHash) Accuracy() float64 {
	return lib.Accuracy(g.hash)
}

func (g GeoHash) ApproximateDistanceTo(other GeoHash) float64 {
	return lib.ApproximateDistance(g.hash, other.hash)
}

func (g GeoHash) ToString() string {
	return g.hash
}

func (g GeoHash) ToLatLong() (float64, float64, error) {
	lat, long, err := lib.ToLatLong(g.hash)
	if err != nil {
		return 0, 0, err
	}
	return lat, long, nil
}
