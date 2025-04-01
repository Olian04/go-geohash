package lib

func ValidateAlphabet(geohash string) bool {
	for _, c := range geohash {
		if _, ok := DecodeMap[c]; !ok {
			return false
		}
	}
	return true
}

func ValidateLength(geohash string) bool {
	length := len(geohash)
	return length >= 0 && length <= 12
}
