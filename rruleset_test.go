package rrule

import (
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: YEARLY, Count: 2, Byweekday: []Weekday{TU},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	r, _ = NewRRule(ROption{Freq: YEARLY, Count: 1, Byweekday: []Weekday{TH},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	value := set.All()
	want := []time.Time{time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC)}
	if !timesEqual(value, want) {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetString(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: YEARLY, Count: 1, Byweekday: []Weekday{TU},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	r, _ = NewRRule(ROption{Freq: YEARLY, Count: 3, Byweekday: []Weekday{TH},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.ExRule(r)
	set.ExDate(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 11, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 18, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC))

	want := `RRULE:FREQ=YEARLY;DTSTART=19970902T090000Z;COUNT=1;BYDAY=TU
RDATE:19970904T090000Z
RDATE:19970909T090000Z
EXRULE:FREQ=YEARLY;DTSTART=19970902T090000Z;COUNT=3;BYDAY=TH
EXDATE:19970904T090000Z
EXDATE:19970911T090000Z
EXDATE:19970918T090000Z`
	value := set.String()
	if want != value {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetRecurrence(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: YEARLY, Count: 1, Byweekday: []Weekday{TU},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	value := set.Recurrence()
	if len(value) != 1 {
		t.Errorf("Wrong length for recurrence got=%v want=%v", len(value), 1)
	}
	want := "RRULE:FREQ=YEARLY;DTSTART=19970902T090000Z;COUNT=1;BYDAY=TU"
	if value[0] != want {
		t.Errorf("get %v, want %v", value[0], want)
	}
}

func TestSetDate(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: YEARLY, Count: 1, Byweekday: []Weekday{TU},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	set.RDate(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC))
	value := set.All()
	want := []time.Time{time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC)}
	if !timesEqual(value, want) {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetExRule(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: YEARLY, Count: 6, Byweekday: []Weekday{TU, TH},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	r, _ = NewRRule(ROption{Freq: YEARLY, Count: 3, Byweekday: []Weekday{TH},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.ExRule(r)
	value := set.All()
	want := []time.Time{time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC)}
	if !timesEqual(value, want) {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetExDate(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: YEARLY, Count: 6, Byweekday: []Weekday{TU, TH},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	set.ExDate(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 11, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 18, 9, 0, 0, 0, time.UTC))
	value := set.All()
	want := []time.Time{time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC)}
	if !timesEqual(value, want) {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetExDateRevOrder(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: MONTHLY, Count: 5, Bymonthday: []int{10},
		Dtstart: time.Date(2004, 1, 1, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	set.ExDate(time.Date(2004, 4, 10, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(2004, 2, 10, 9, 0, 0, 0, time.UTC))
	value := set.All()
	want := []time.Time{time.Date(2004, 1, 10, 9, 0, 0, 0, time.UTC),
		time.Date(2004, 3, 10, 9, 0, 0, 0, time.UTC),
		time.Date(2004, 5, 10, 9, 0, 0, 0, time.UTC)}
	if !timesEqual(value, want) {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetDateAndExDate(t *testing.T) {
	set := Set{}
	set.RDate(time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 11, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 18, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 11, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 18, 9, 0, 0, 0, time.UTC))
	value := set.All()
	want := []time.Time{time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC)}
	if !timesEqual(value, want) {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetDateAndExRule(t *testing.T) {
	set := Set{}
	set.RDate(time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 11, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 18, 9, 0, 0, 0, time.UTC))
	r, _ := NewRRule(ROption{Freq: YEARLY, Count: 3, Byweekday: []Weekday{TH},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.ExRule(r)
	value := set.All()
	want := []time.Time{time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 16, 9, 0, 0, 0, time.UTC)}
	if !timesEqual(value, want) {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetBefore(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: DAILY, Count: 7,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	want := time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC)
	value := set.Before(time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC), false)
	if value != want {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetBeforeInc(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: DAILY, Count: 7,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	want := time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC)
	value := set.Before(time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC), true)
	if value != want {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetAfter(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: DAILY, Count: 7,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	want := time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC)
	value := set.After(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC), false)
	if value != want {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetAfterInc(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: DAILY, Count: 7,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	want := time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC)
	value := set.After(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC), true)
	if value != want {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetBetween(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: DAILY, Count: 7,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	value := set.Between(time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC), time.Date(1997, 9, 6, 9, 0, 0, 0, time.UTC), false)
	want := []time.Time{time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC)}
	if !timesEqual(value, want) {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetBetweenInc(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: DAILY, Count: 7,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	value := set.Between(time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC), time.Date(1997, 9, 6, 9, 0, 0, 0, time.UTC), true)
	want := []time.Time{time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 6, 9, 0, 0, 0, time.UTC)}
	if !timesEqual(value, want) {
		t.Errorf("get %v, want %v", value, want)
	}
}
