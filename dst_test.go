package rrule

import (
	"testing"
	"time"
)

func TestDST_HourlyDSTStart(t *testing.T) {
	sydLoc, _ := time.LoadLocation("Australia/Sydney")
	r, _ := NewRRule(ROption{Freq: HOURLY, Interval: 1, Count: 3,
		Dtstart: time.Date(2022, 10, 2, 1, 0, 0, 0, sydLoc),
	})
	got := r.All()
	want := []string{
		"2022-10-02 01:00:00 +1000 AEST",
		"2022-10-02 03:00:00 +1100 AEDT",
		"2022-10-02 04:00:00 +1100 AEDT",
	}
	for i, g := range got {
		if g.String() != want[i] {
			t.Errorf("got: %v, want: %v", g, want[i])
		}
	}
	var utcTimes []time.Time
	for _, dt := range got {
		utcTimes = append(utcTimes, dt.UTC())
	}
	want = []string{
		"2022-10-01 15:00:00 +0000 UTC",
		"2022-10-01 16:00:00 +0000 UTC",
		"2022-10-01 17:00:00 +0000 UTC",
	}

	for i, g := range utcTimes {
		if g.String() != want[i] {
			t.Errorf("got: %v, want: %v", g, want[i])
		}
	}
}

func TestDST_HourlyDSTEnd(t *testing.T) {
	sydLoc, _ := time.LoadLocation("Australia/Sydney")
	r, _ := NewRRule(ROption{Freq: HOURLY, Interval: 1, Count: 3,
		Dtstart: time.Date(2023, 4, 2, 1, 0, 0, 0, sydLoc),
	})
	got := r.All()
	want := []string{
		"2023-04-02 01:00:00 +1100 AEDT",
		"2023-04-02 02:00:00 +1100 AEDT",
		"2023-04-02 02:00:00 +1000 AEST",
	}
	for i, g := range got {
		if g.String() != want[i] {
			t.Errorf("got: %v, want: %v", g, want[i])
		}
	}

	var utcTimes []time.Time
	for _, dt := range got {
		utcTimes = append(utcTimes, dt.UTC())
	}
	want = []string{
		"2023-04-01 14:00:00 +0000 UTC",
		"2023-04-01 15:00:00 +0000 UTC",
		"2023-04-01 16:00:00 +0000 UTC",
	}
	for i, g := range utcTimes {
		if g.String() != want[i] {
			t.Errorf("got: %v, want: %v", g, want[i])
		}
	}
}
