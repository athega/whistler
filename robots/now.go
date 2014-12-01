package robots

import "time"

type now struct {
	time.Time
}

// Now creates an instance of now
func Now(t time.Time) *now {
	return &now{t}
}

func (now *now) BeginningOfHour() time.Time {
	return now.Truncate(time.Hour)
}

func (now *now) BeginningOfDay() time.Time {
	d := time.Duration(-now.Hour()) * time.Hour

	return now.BeginningOfHour().Add(d)
}

func (now *now) EndOfDay() time.Time {
	return now.BeginningOfDay().Add(24*time.Hour - time.Nanosecond)
}

func (now *now) Friday() time.Time {
	t := now.BeginningOfDay()

	var d int

	switch t.Weekday() {
	case time.Monday:
		d = 4
	case time.Tuesday:
		d = 3
	case time.Wednesday:
		d = 2
	case time.Thursday:
		d = 1
	case time.Friday:
		d = 0
	case time.Saturday:
		d = 6
	case time.Sunday:
		d = 5
	}

	return t.Truncate(time.Hour).
		Add(time.Duration(d) * 24 * time.Hour)
}
