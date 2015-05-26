package gigasecond

import (
	"math"
	"time"
)

const TestVersion = 2

var Birthday = time.Date(1977, time.November, 10, 23, 0, 0, 0, time.UTC)

const oneGigaSecond = 1000000000 * time.Second

func AddGigasecond(time time.Time) time.Time {
	return time.Add(oneGigaSecond)
}
