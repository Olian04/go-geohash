package lib

import "fmt"

var BITS = []int{16, 8, 4, 2, 1}

func FromLatLong(latitude float64, longitude float64) string {
	isEven := true
	lat := make([]float64, 2)
	lon := make([]float64, 2)
	bit := 0
	ch := 0
	precision := 12
	geohash := ""

	lat[0] = -90.0
	lat[1] = 90.0
	lon[0] = -180.0
	lon[1] = 180.0

	for len(geohash) < precision {
		if isEven {
			mid := (lon[0] + lon[1]) / 2
			if longitude > mid {
				ch |= BITS[bit]
				lon[0] = mid
			} else {
				lon[1] = mid
			}
		} else {
			mid := (lat[0] + lat[1]) / 2
			if latitude > mid {
				ch |= BITS[bit]
				lat[0] = mid
			} else {
				lat[1] = mid
			}
		}

		isEven = !isEven
		if bit < 4 {
			bit++
		} else {
			geohash += string(EncodeMap[ch])
			bit = 0
			ch = 0
		}
	}
	return geohash
}

func ToLatLong(geohash string) (float64, float64, error) {
	if !ValidateLength(geohash) {
		return 0.0, 0.0, fmt.Errorf("geohash: %s is not of valid length", geohash)
	}
	if !ValidateAlphabet(geohash) {
		return 0.0, 0.0, fmt.Errorf("geohash: %s contains invalid characters", geohash)
	}

	isEven := true
	lat := make([]float64, 3)
	lon := make([]float64, 3)
	lat[0] = -90.0
	lat[1] = 90.0
	lon[0] = -180.0
	lon[1] = 180.0
	latErr := 90.0
	lonErr := 180.0

	for i := 0; i < len(geohash); i++ {
		c := geohash[i]
		cd := DecodeMap[rune(c)]
		for j := 0; j < 5; j++ {
			mask := BITS[j]
			if isEven {
				lonErr /= 2
				refineInterval(lon, cd, mask)
			} else {
				latErr /= 2
				refineInterval(lat, cd, mask)
			}
			isEven = !isEven
		}
	}

	lat[2] = (lat[0] + lat[1]) / 2
	lon[2] = (lon[0] + lon[1]) / 2

	return lat[2], lon[2], nil
}

func refineInterval(interval []float64, cd int, mask int) {
	if cd&mask != 0 {
		interval[0] = (interval[0] + interval[1]) / 2
	} else {
		interval[1] = (interval[0] + interval[1]) / 2
	}
}
