package clock

import "fmt"

const TestVersion = 2

type Clock struct {
	hour   int
	minute int
}

func Time(hour int, minute int) Clock {
	// Convert negative minutes into negative hours
	for minute < 0 {
		minute += 60
		hour -= 1
	}

	// Convert negative hours by stepping forward by days
	for hour < 0 {
		hour += 24
	}

	// Convert overflowing minutes into positive hours
	hour = hour + minute/60
	minute = minute % 60

	// Cut off overflowing hours
	hour = hour % 24

	// ...and we are good.
	return Clock{hour, minute}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hour, clock.minute)
}

func (clock Clock) Add(minutes int) Clock {
	return Time(clock.hour, clock.minute+minutes)
}
