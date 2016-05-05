package bob

import (
	"strings"
	"unicode"
)

// testVersion should match the targetTestVersion in the test file.
const testVersion = 2

// Hey responds to a greeting like a teenager.
func Hey(greeting string) (answer string) {

	greeting = strings.TrimSpace(greeting)

	switch {

		case greeting == "":
			answer = "Fine. Be that way!"

		case yelled(greeting, unicode.IsUpper) && !yelled(greeting, unicode.IsLower):
			answer = "Whoa, chill out!"

		case greeting[len(greeting)-1] == '?':
			answer = "Sure."

		default:
			answer = "Whatever."
	}

	return

}

// yelled checks if you yell at him.
func yelled(words string, f func(rune) bool) bool {

	for _, word := range words {
		if f(word) {
			return true
		}
	}

	return false

}