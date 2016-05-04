package hamming

import (
	"errors"
)

// testVersion should match the targetTestVersion in the test file.
const testVersion = 4

// Distance calculates the Hamming difference between two DNA strands
func Distance(a, b string) (distance int, err error) {

	if len(a) != len(b) {
		return distance, errors.New("Wrong lengths")
	}

	for i := 0; i<len(a) && i<len(b); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return

}