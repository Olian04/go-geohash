package internal

import (
	"testing"
)

func TestApproximateDistanceForSameHash(t *testing.T) {
	distance := ApproximateDistance("9q8y7x6w5v4u", "9q8y7x6w5v4u")
	expected := 0.02
	if distance != expected {
		t.Errorf("Distance should be %f but got %f", expected, distance)
	}
}

func TestApproximateDistanceForDifferentHashesOfSameLength(t *testing.T) {
	distance := ApproximateDistance("9q8y7x6w5v4u", "9q8y7x6w5v4t")
	expected := 0.09
	if distance != expected {
		t.Errorf("Distance should be %f but got %f", expected, distance)
	}
}

func TestApproximateDistanceForHashesOfDifferentLength(t *testing.T) {
	distance := ApproximateDistance("9q8y7x6w5v4u", "9q8y7")
	expected := 3803.0
	if distance != expected {
		t.Errorf("Distance should be %f but got %f", expected, distance)
	}
}

func TestApproximateDistanceAgainstEmptyHash(t *testing.T) {
	distance := ApproximateDistance("9q8y7x6w5v4u", "")
	expected := 20_000_000.0
	if distance != expected {
		t.Errorf("Distance should be %f but got %f", expected, distance)
	}
}

func TestAccuracy(t *testing.T) {
	accuracy := Accuracy("9q8y7")
	expected := 3803.0
	if accuracy != expected {
		t.Errorf("Accuracy should be %f but got %f", expected, accuracy)
	}
	accuracy = Accuracy("9q8y7x6w5v4u")
	expected = 0.02
	if accuracy != expected {
		t.Errorf("Accuracy should be %f but got %f", expected, accuracy)
	}
}
