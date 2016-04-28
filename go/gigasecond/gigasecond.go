package gigasecond

import "time"

// testVersion should match the targetTestVersion in the test file.
const testVersion = 4

// A gigasecond is one billion (1e9) seconds.
const Gigasecond = time.Duration(1e9) * time.Second

// AddGigasecond adds one gigasecond to a given time.Time.
func AddGigasecond(date time.Time) time.Time {
	return date.Add(Gigasecond)
}

// Birthday is my birthday and is used to celebrate my Gs anniversary.
var Birthday, _ = time.Parse("2006-01-02", "1977-03-10")