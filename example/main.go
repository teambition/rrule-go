package main

import (
	"fmt"
	"time"

	"github.com/teambition/rrule-go"
)

func exampleRRule() {
	fmt.Println("Daily, for 10 occurrences.")
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   10,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	fmt.Println(r.All())
	// [1997-09-02 09:00:00 +0800 CST
	// 1997-09-03 09:00:00 +0800 CST
	// ...
	// 1997-09-10 09:00:00 +0800 CST
	// 1997-09-11 09:00:00 +0800 CST]

	fmt.Println(r.Between(
		time.Date(1997, 9, 6, 0, 0, 0, 0, time.Local),
		time.Date(1997, 9, 8, 0, 0, 0, 0, time.Local), true))
	// [1997-09-06 09:00:00 +0800 CST
	// 1997-09-07 09:00:00 +0800 CST]

	fmt.Println(r)
	// FREQ=DAILY;DTSTART=19970902T010000Z;COUNT=10

	fmt.Println("\nEvery four years, the first Tuesday after a Monday in November, 3 occurrences (U.S. Presidential Election day).")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:       rrule.YEARLY,
		Interval:   4,
		Count:      3,
		Bymonth:    []int{11},
		Byweekday:  []rrule.Weekday{rrule.TU},
		Bymonthday: []int{2, 3, 4, 5, 6, 7, 8},
		Dtstart:    time.Date(1996, 11, 5, 9, 0, 0, 0, time.Local)})
	fmt.Println(r.All())
	// [1996-11-05 09:00:00 +0800 CST
	// 2000-11-07 09:00:00 +0800 CST
	// 2004-11-02 09:00:00 +0800 CST]
}

func exampleRRuleSet() {
	fmt.Println("\nDaily, for 7 days, jumping Saturday and Sunday occurrences.")
	set := rrule.Set{}
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   7,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	set.RRule(r)
	fmt.Println(set.All())
	// [1997-09-02 09:00:00 +0800 CST
	// 1997-09-03 09:00:00 +0800 CST
	// 1997-09-04 09:00:00 +0800 CST
	// 1997-09-05 09:00:00 +0800 CST
	// 1997-09-08 09:00:00 +0800 CST]

	fmt.Println("\nWeekly, for 4 weeks, plus one time on day 7, and not on day 16.")
	set = rrule.Set{}
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:    rrule.WEEKLY,
		Count:   4,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	set.RRule(r)
	set.RDate(time.Date(1997, 9, 7, 9, 0, 0, 0, time.Local))
	set.ExDate(time.Date(1997, 9, 16, 9, 0, 0, 0, time.Local))
	fmt.Println(set.All())
	// [1997-09-02 09:00:00 +0800 CST
	// 1997-09-07 09:00:00 +0800 CST
	// 1997-09-09 09:00:00 +0800 CST
	// 1997-09-23 09:00:00 +0800 CST]
}

func exampleStrToRRule() {
	fmt.Println()
	r, _ := rrule.StrToRRule("FREQ=DAILY;INTERVAL=10;COUNT=5")
	fmt.Println(r.All())
	// [2017-03-15 14:12:02 +0800 CST
	// 2017-03-25 14:12:02 +0800 CST
	// 2017-04-04 14:12:02 +0800 CST
	// 2017-04-14 14:12:02 +0800 CST
	// 2017-04-24 14:12:02 +0800 CST]
}

func exampleStrToRRuleSet() {
	s, _ := rrule.StrToRRuleSet(`RRULE:FREQ=DAILY;INTERVAL=10;COUNT=5
RDATE:20060102T150405Z`)
	fmt.Println(s.All())
}

func main() {
	exampleRRule()
	exampleRRuleSet()
	exampleStrToRRule()
	exampleStrToRRuleSet()
}
