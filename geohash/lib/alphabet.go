package lib

import "fmt"

const (
	Base32ghs = "0123456789bcdefghjkmnpqrstuvwxyz"
)

var DecodeMap = map[rune]int{}
var EncodeMap = map[int]rune{}

func init() {
	for i, c := range Base32ghs {
		DecodeMap[c] = i
		EncodeMap[i] = c
	}
}

// Decodes a rune from the Base32ghs alphabet to an integer
func Decode(c rune) (int, error) {
	i, ok := DecodeMap[c]
	if !ok {
		return 0, fmt.Errorf("invalid character: %c", c)
	}
	return i, nil
}

// Encodes an integer to a rune in the Base32ghs alphabet
func Encode(i int) (rune, error) {
	r, ok := EncodeMap[i]
	if !ok {
		return 0, fmt.Errorf("invalid integer: %d", i)
	}
	return r, nil
}
