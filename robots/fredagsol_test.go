package robots

import (
	"testing"
	"time"
)

var isItFredagsölTests = []struct {
	in   time.Time
	ok   bool
	text string
}{
	{parseTime("Fri, 21 Nov 2014 17:04:05 CET"), true, "Fredagsöl is now! :beer: :beers:"},
	{parseTime("Fri, 21 Nov 2014 22:14:33 CET"), true, "Fredagsöl is now! :beer: :beers:"},
	{parseTime("Sat, 22 Nov 2014 10:24:10 CET"), false, "Next Fredagsöl: 150h35m50s"},
	{parseTime("Wed, 26 Nov 2014 23:05:00 CET"), false, "Next Fredagsöl: 41h55m0s"},
	{parseTime("Thu, 27 Nov 2014 17:00:00 CET"), false, "Next Fredagsöl: 24h0m0s"},
	{parseTime("Fri, 28 Nov 2014 16:11:28 CET"), false, "Next Fredagsöl: 48m32s"},
	{parseTime("Fri, 28 Nov 2014 16:59:43 CET"), false, "Next Fredagsöl: 17s"},
}

func TestIsItFredagsöl(t *testing.T) {
	b := NewFredagsölBot()

	for _, tt := range isItFredagsölTests {
		if text, ok := b.isItFredagsöl(tt.in); text != tt.text || ok != tt.ok {
			t.Errorf(`b.isItFredagsöl(%s) = "%v", %v, want "%v", %v`,
				tt.in, text, ok, tt.text, tt.ok)
		}
	}
}

func parseTime(value string) time.Time {
	t, _ := time.Parse(time.RFC1123, value)
	return t
}
