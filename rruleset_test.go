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
	value := set.All()
	want := []time.Time{time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC)}
	if !timesEqual(value, want) {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetOverlapping(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: YEARLY,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	v1 := set.All()
	if len(v1) > 300 || len(v1) < 200 {
		t.Errorf("No default Util time")
	}
}

func TestSetString(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: YEARLY, Count: 1, Byweekday: []Weekday{TU},
		Dtstart: time.Date(1997, 9, 2, 8, 0, 0, 0, time.UTC)})
	set.RRule(r)
	set.ExDate(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 11, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 18, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC))

	want := `DTSTART:19970902T080000Z
RRULE:FREQ=YEARLY;COUNT=1;BYDAY=TU
RDATE:19970904T090000Z
RDATE:19970909T090000Z
EXDATE:19970904T090000Z
EXDATE:19970911T090000Z
EXDATE:19970918T090000Z`
	value := set.String()
	if want != value {
		t.Errorf("get %v, want %v", value, want)
	}
}

func TestSetDTStart(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: YEARLY, Count: 1, Byweekday: []Weekday{TU},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	set.ExDate(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 11, 9, 0, 0, 0, time.UTC))
	set.ExDate(time.Date(1997, 9, 18, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC))
	set.RDate(time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC))

	nyLoc, _ := time.LoadLocation("America/New_York")
	set.DTStart(time.Date(1997, 9, 3, 9, 0, 0, 0, nyLoc))

	want := `DTSTART;TZID=America/New_York:19970903T090000
RRULE:FREQ=YEARLY;COUNT=1;BYDAY=TU
RDATE:19970904T090000Z
RDATE:19970909T090000Z
EXDATE:19970904T090000Z
EXDATE:19970911T090000Z
EXDATE:19970918T090000Z`
	value := set.String()
	if want != value {
		t.Errorf("get \n%v\n want \n%v\n", value, want)
	}

	sset, err := StrToRRuleSet(set.String())
	if err != nil {
		t.Errorf("Could not create RSET from set output")
	}
	if sset.String() != set.String() {
		t.Errorf("RSET created from set output different than original set, %s", sset.String())
	}
}

func TestSetRecurrence(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: YEARLY, Count: 1, Byweekday: []Weekday{TU},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	value := set.Recurrence()
	if len(value) != 2 {
		t.Errorf("Wrong length for recurrence got=%v want=%v", len(value), 2)
	}
	want := "DTSTART:19970902T090000Z\nRRULE:FREQ=YEARLY;COUNT=1;BYDAY=TU"
	if set.String() != want {
		t.Errorf("get %s, want %v", set.String(), want)
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

func TestSetRDates(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: YEARLY, Count: 1, Byweekday: []Weekday{TU},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	set.SetRDates([]time.Time{
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
	})
	value := set.All()
	want := []time.Time{
		time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 9, 9, 0, 0, 0, time.UTC),
	}
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

func TestSetExDates(t *testing.T) {
	set := Set{}
	r, _ := NewRRule(ROption{Freq: YEARLY, Count: 6, Byweekday: []Weekday{TU, TH},
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.UTC)})
	set.RRule(r)
	set.SetExDates([]time.Time{
		time.Date(1997, 9, 4, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 11, 9, 0, 0, 0, time.UTC),
		time.Date(1997, 9, 18, 9, 0, 0, 0, time.UTC),
	})
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

func TestSetTrickyTimeZones(t *testing.T) {
	set := Set{}

	moscow, _ := time.LoadLocation("Europe/Moscow")
	newYork, _ := time.LoadLocation("America/New_York")
	tehran, _ := time.LoadLocation("Asia/Tehran")

	r, _ := NewRRule(ROption{
		Freq:    DAILY,
		Count:   4,
		Dtstart: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).In(moscow),
	})
	set.RRule(r)

	set.ExDate(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).In(newYork))
	set.ExDate(time.Date(2000, 1, 2, 0, 0, 0, 0, time.UTC).In(tehran))
	set.ExDate(time.Date(2000, 1, 3, 0, 0, 0, 0, time.UTC).In(moscow))
	set.ExDate(time.Date(2000, 1, 4, 0, 0, 0, 0, time.UTC))

	occurrences := set.All()

	if len(occurrences) > 0 {
		t.Errorf("No all occurrences excluded by ExDate: [%+v]", occurrences)
	}
}

func TestSetDtStart(t *testing.T) {
	ogr := []string{"DTSTART;TZID=America/Los_Angeles:20181115T000000", "RRULE:FREQ=DAILY;INTERVAL=1;WKST=SU;UNTIL=20181117T235959"}
	set, _ := StrSliceToRRuleSet(ogr)

	ogoc := set.All()
	set.DTStart(set.GetDTStart().AddDate(0, 0, 1))

	noc := set.All()
	if len(noc) != len(ogoc)-1 {
		t.Fatalf("As per the new DTStart the new occurences should exactly be one less that the original, new :%d original: %d", len(noc), len(ogoc))
	}

	for i := range noc {
		if noc[i] != ogoc[i+1] {
			t.Errorf("New occurences should just offset by one, mismatch at %d, expected: %+v, actual: %+v", i, ogoc[i+1], noc[i])
		}
	}
}

func TestRuleSetChangeDTStartTimezoneRespected(t *testing.T) {
	/*
		https://golang.org/pkg/time/#LoadLocation

		"The time zone database needed by LoadLocation may not be present on all systems, especially non-Unix systems.
		LoadLocation looks in the directory or uncompressed zip file named by the ZONEINFO environment variable,
		if any, then looks in known installation locations on Unix systems, and finally looks in
		$GOROOT/lib/time/zoneinfo.zip."
	*/
	loc, err := time.LoadLocation("CET")
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}

	ruleSet := &Set{}
	rule, err := NewRRule(
		ROption{
			Freq:     DAILY,
			Count:    10,
			Wkst:     MO,
			Byhour:   []int{10},
			Byminute: []int{0},
			Bysecond: []int{0},
			Dtstart:  time.Date(2019, 3, 6, 0, 0, 0, 0, loc),
		},
	)
	if err != nil {
		t.Fatal("expected", nil, "got", err)
	}
	ruleSet.RRule(rule)
	ruleSet.DTStart(time.Date(2019, 3, 6, 0, 0, 0, 0, time.UTC))

	events := ruleSet.All()
	if len(events) != 10 {
		t.Fatal("expected", 10, "got", len(events))
	}

	for _, e := range events {
		if e.Location().String() != "UTC" {
			t.Fatal("expected", "UTC", "got", e.Location().String())
		}
	}
}
