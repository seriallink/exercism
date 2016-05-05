package raindrops

import "fmt"

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// Convert converts a number to raindrop sounds
func Convert(number int) (sounds string) {

	if number%3 == 0 {
		sounds += "Pling"
	}

	if number%5 == 0 {
		sounds += "Plang"
	}

	if number%7 == 0 {
		sounds += "Plong"
	}

	if sounds == "" {
		sounds = fmt.Sprintf("%d",number)
	}

	return

}