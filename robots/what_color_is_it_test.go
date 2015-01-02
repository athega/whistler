package robots

import (
	"testing"
	"time"
)

var colorHexTests = []struct {
	time time.Time
	hex  string
}{
	{parseTime("Fri, 21 Nov 2014 17:04:05 CET"), "#170405"},
	{parseTime("Sat, 22 Nov 2014 10:24:10 CET"), "#102410"},
	{parseTime("Wed, 26 Nov 2014 23:05:00 CET"), "#230500"},
	{parseTime("Fri, 28 Nov 2014 16:59:43 CET"), "#165943"},
}

func TestColorHex(t *testing.T) {
	b := NewWhatColorIsItBot()

	for _, tt := range colorHexTests {
		if got := b.colorHex(tt.time); got != tt.hex {
			t.Errorf("b.colorHex(%#v) = %v, want %v", tt.time, got, tt.hex)
		}
	}
}
