package clock

import "fmt"

// testVersion should match the targetTestVersion in the test file.
const testVersion = 4

// Clock manages time in minutes.
type Clock int

// Create a new Clock.
func New(hour, minute int) Clock {

	minutes := (hour*60 + minute) % (60 * 24)

	if minutes < 0 {
		minutes += 60 * 24
	}

	return Clock(minutes)

}

// String converts Clock in hh:mm format.
func (this Clock) String() string {
	return fmt.Sprintf("%02d:%02d", this/60, this%60)
}

// Add creates a new Clock adding minutes to it.
func (this Clock) Add(minutes int) Clock {
	return New(0, int(this)+minutes)
}