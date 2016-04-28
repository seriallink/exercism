package hello

import "fmt"

// Defines the version of the unit test that this will pass.
const testVersion = 2

// HelloWorld greets you or the world if you neglect to give it your name.
func HelloWorld(name string) string {

	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("Hello, %s!", name)

}