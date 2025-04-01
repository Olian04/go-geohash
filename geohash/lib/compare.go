package lib

var distancePerCharacterInMeters = map[int]float64{
	0:  20_000_000,
	1:  5_003_530,
	2:  625_441,
	3:  123_264,
	4:  19_545,
	5:  3_803,
	6:  610,
	7:  118,
	8:  19,
	9:  3.71,
	10: 0.6,
	11: 0.09,
	12: 0.02,
}

// ApproximateDistance returns the distance between two geohashes in meters.
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
func ApproximateDistance(hash1 string, hash2 string) float64 {
	matchingLength := equalityLength(hash1, hash2)
	return distancePerCharacterInMeters[matchingLength]
}

func EqualityScore(hash1 string, hash2 string) int {
	return equalityLength(hash1, hash2)
}

// Accuracy returns the accuracy of the geohash in meters
func Accuracy(hash string) float64 {
	return distancePerCharacterInMeters[len(hash)]
}

// equalityLength returns the length of the subset of the two hashes that are the same
func equalityLength(hash1 string, hash2 string) int {
	minLength := min(len(hash1), len(hash2))
	matchingLength := 0
	for i := 0; i < minLength; i++ {
		if hash1[i] != hash2[i] {
			break
		}
		matchingLength++
	}
	return matchingLength
}
