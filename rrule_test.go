package rrule

import (
	"testing"
	"time"
)

func TestNewRRule(t *testing.T) {
	r, e := NewRRule(ROption{Freq: MONTHLY})
	if seconds := time.Now().Sub(r.dtstart).Seconds(); seconds > 10 {
		t.Errorf(`time.Now().Sub(r.dtstrt).Seconds() = %f, want < 10`, seconds)
	}

	r, e = NewRRule(ROption{Bysetpos: []int{367}})
	if e == nil {
		t.Errorf("_, e = NewRRule(ROption{Bysetpos:[]int{367}}); e = nil, want error")
	}

	r, _ = NewRRule(ROption{
		Freq:    WEEKLY,
		Dtstart: time.Date(2017, 3, 14, 9, 0, 0, 0, time.UTC),
	})
	if len(r.byweekday) != 1 || r.byweekday[0] != TU.weekday {
		t.Errorf("r.byweekday = %v, want [1]", r.byweekday)
	}

	r, _ = NewRRule(ROption{
		Bymonthday: []int{-1},
		Byweekday:  []Weekday{MO.Nth(1)},
	})
	if len(r.bynmonthday) != 1 || r.bynmonthday[0] != -1 {
		t.Errorf("r.bynmonthday = %v, want [-1]", r.bynmonthday)
	}
	if len(r.bynweekday) != 1 || r.bynweekday[0] != MO.Nth(1) {
		t.Errorf("r.bynweekday = %v, want +1MO", r.bynweekday)
	}
}

func TestRRule(t *testing.T) {
	r, _ := NewRRule(ROption{
		Freq:    YEARLY,
		Count:   4,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC),
	})
	instances := r.All()
	if len := len(instances); len != 4 {
		t.Errorf("len(r.All()) = %d, want 4", len)
	}
	for i, year := range []int{1997, 1998, 1999, 2000} {
		want := time.Date(year, 9, 2, 9, 0, 0, 0, time.UTC)
		if instances[i] != want {
			t.Errorf("r.All()[i] = %v, want %v", instances[i], want)
		}
	}
	from, to := time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC), time.Date(2000, 9, 2, 9, 0, 0, 0, time.UTC)
	if instances := r.Between(from, to, true); len(instances) != 4 {
		t.Errorf("len(r.Between(%v, %v, true)) = %d, want 4", from, to, len(instances))
	}
	if instances := r.Between(from, to, false); len(instances) != 2 {
		t.Errorf("len(r.Between(%v, %v, false)) = %d, want 2", from, to, len(instances))
	}
	if instance := r.Before(from, true); instance != from {
		t.Errorf("r.Before(%v, true) = %v, want %v", from, instance, from)
	}
	if instance := r.After(to, true); instance != to {
		t.Errorf("r.After(%v, true) = %v, want %v", to, instance, to)
	}
	if instance := r.After(to, false); !instance.IsZero() {
		t.Errorf("r.After(%v, false) = %v, want %v", to, instance, time.Time{})
	}
}

func TestRRuleIterate(t *testing.T) {
	r, _ := NewRRule(ROption{
		Freq:      YEARLY,
		Dtstart:   time.Date(1997, 9, 2, 0, 0, 0, 0, time.UTC),
		Byweekday: []Weekday{FR.Nth(1), SA.Nth(-1)},
		Byweekno:  []int{1, 2, -100},
		Bymonth:   []int{1},
	})
	if value, want := r.All()[0], time.Date(1998, 1, 2, 0, 0, 0, 0, time.UTC); value != want {
		t.Errorf("r.All()[0] = %v, want %v", value, want)
	}

	r, _ = NewRRule(ROption{
		Freq:     YEARLY,
		Dtstart:  time.Date(1997, 9, 2, 0, 0, 0, 0, time.UTC),
		Byeaster: []int{1},
		Until:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	})
	if value, want := r.All()[0], time.Date(1998, 4, 12, 0, 0, 0, 0, time.UTC); value != want {
		t.Errorf("r.All()[0] = %v, want %v", value, want)
	}

	r, _ = NewRRule(ROption{
		Freq:      YEARLY,
		Dtstart:   time.Date(1997, 9, 2, 0, 0, 0, 0, time.UTC),
		Byweekno:  []int{-1},
		Byweekday: []Weekday{MO.Nth(-1)},
		Count:     3,
	})
	if value, want := r.All()[0], time.Date(1998, 12, 28, 0, 0, 0, 0, time.UTC); value != want {
		t.Errorf("r.All()[0] = %v, want %v", value, want)
	}

	r, _ = NewRRule(ROption{
		Freq:      MONTHLY,
		Interval:  15,
		Dtstart:   time.Date(1997, 9, 2, 0, 0, 0, 0, time.UTC),
		Byweekday: []Weekday{MO.Nth(1)},
	})
	if value, want := r.All()[1], time.Date(2000, 3, 6, 0, 0, 0, 0, time.UTC); value != want {
		t.Errorf("r.All()[1] = %v, want %v", value, want)
	}

	r, _ = NewRRule(ROption{
		Freq:    WEEKLY,
		Dtstart: time.Date(1997, 9, 2, 0, 0, 0, 0, time.UTC),
	})
	if value, want := r.All()[1], time.Date(1997, 9, 9, 0, 0, 0, 0, time.UTC); value != want {
		t.Errorf("r.All()[1] = %v, want %v", value, want)
	}

	r, _ = NewRRule(ROption{
		Freq:    WEEKLY,
		Wkst:    TU,
		Dtstart: time.Date(1997, 9, 1, 0, 0, 0, 0, time.UTC),
	})
	if value, want := r.All()[1], time.Date(1997, 9, 8, 0, 0, 0, 0, time.UTC); value != want {
		t.Errorf("r.All()[1] = %v, want %v", value, want)
	}

	r, _ = NewRRule(ROption{
		Freq:       HOURLY,
		Dtstart:    time.Date(1997, 9, 1, 23, 0, 0, 0, time.UTC),
		Until:      time.Date(1997, 9, 3, 11, 0, 0, 0, time.UTC),
		Bysetpos:   []int{1, -1, 61},
		Bymonthday: []int{2, 3},
		Byhour:     []int{0, 2},
		Byminute:   []int{0, 1, 2},
	})
	if value, want := r.All()[3], time.Date(1997, 9, 2, 2, 2, 0, 0, time.UTC); value != want {
		t.Errorf("r.All()[3] = %v, want %v", value, want)
	}

	r, _ = NewRRule(ROption{
		Freq:       MINUTELY,
		Dtstart:    time.Date(1997, 9, 2, 23, 0, 0, 0, time.UTC),
		Count:      4,
		Bymonthday: []int{2, 4},
		Byminute:   []int{0},
		Bysecond:   []int{0},
	})
	if value, want := r.All()[3], time.Date(1997, 9, 4, 2, 0, 0, 0, time.UTC); value != want {
		t.Errorf("r.All()[3] = %v, want %v", value, want)
	}

	r, _ = NewRRule(ROption{
		Freq:       SECONDLY,
		Dtstart:    time.Date(1997, 12, 31, 23, 59, 59, 0, time.UTC),
		Count:      4,
		Bysetpos:   []int{1, -1},
		Bymonthday: []int{31, 2},
	})
	if value, want := r.All()[3], time.Date(1998, 1, 2, 0, 0, 2, 0, time.UTC); value != want {
		t.Errorf("r.All()[3] = %v, want %v", value, want)
	}
}

func TestRRuleSet(t *testing.T) {
	// Daily, for 7 days, jumping Saturday and Sunday occurrences.
	set := Set{}
	r, _ := NewRRule(ROption{
		Freq:    DAILY,
		Count:   7,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	r, _ = NewRRule(ROption{
		Freq:      YEARLY,
		Byweekday: []Weekday{SA, SU},
		Dtstart:   time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.ExRule(r)
	set.RDate(time.Date(1997, 9, 7, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 3, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC))
	// 2,4,5,8
	if value, want := set.All()[2], time.Date(1997, 9, 8, 9, 0, 0, 0, time.UTC); value != want {
		t.Errorf("set.All()[4] = %v, want %v", value, want)
	}
	from, to := time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC), time.Date(1997, 9, 5, 9, 0, 0, 0, time.UTC)
	if len := len(set.Between(from, to, true)); len != 2 {
		t.Errorf("len(set.Between(%v, %v, true)) = %d, want 2", from, to, len)
	}
	if instance := set.Before(from, true); instance != from {
		t.Errorf("set.Before(%v, true) = %v, want %v", from, instance, from)
	}
	if instance := set.After(to, true); instance != to {
		t.Errorf("set.After(%v, true) = %v, want %v", to, instance, to)
	}
}

func TestStr(t *testing.T) {
	str := "FREQ=WEEKLY;DTSTART=20120201T093000Z;INTERVAL=5;WKST=TU;COUNT=2;UNTIL=20130130T230000Z;BYSETPOS=2;BYMONTH=3;BYYEARDAY=95;BYWEEKNO=1;BYDAY=MO,+2FR;BYHOUR=9;BYMINUTE=30;BYSECOND=0;BYEASTER=-1"
	r, _ := StrToRRule(str)
	if s := r.String(); s != str {
		t.Errorf("StrToRRule(%q).String() = %q, want %q", str, s, str)
	}
}

func TestInvalidString(t *testing.T) {
	cases := []string{
		"",
		"FREQ",
		"FREQ=HELLO",
		"BYMONTH=",
		"FREQ=WEEKLY;HELLO=WORLD",
		"FREQ=WEEKLY;BYMONTHDAY=I",
		"FREQ=WEEKLY;BYDAY=M",
		"FREQ=WEEKLY;BYDAY=MQ",
		"FREQ=WEEKLY;BYDAY=+MO",
	}
	for _, item := range cases {
		if _, e := StrToRRule(item); e == nil {
			t.Errorf("StrToRRule(%q) = nil, want error", item)
		}
	}
}
