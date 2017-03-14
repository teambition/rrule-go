package main

import (
	"fmt"
	"time"

	"github.com/teambition/rrule-go"
)

func testRRule() {
	fmt.Println("Daily, for 10 occurrences.")
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   10,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nDaily until December 24, 1997")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local),
		Until:   time.Date(1997, 12, 24, 0, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery other day, 5 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:     rrule.DAILY,
		Interval: 2,
		Count:    5,
		Dtstart:  time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery 10 days, 5 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:     rrule.DAILY,
		Interval: 10,
		Count:    5,
		Dtstart:  time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEveryday in January, for 3 years.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.YEARLY,
		Bymonth:   []int{int(time.January)},
		Byweekday: []rrule.Weekday{rrule.MO, rrule.TU, rrule.WE, rrule.TH, rrule.FR, rrule.SA, rrule.SU},
		Dtstart:   time.Date(1998, 1, 1, 9, 0, 0, 0, time.Local),
		Until:     time.Date(2000, 1, 31, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nSame thing, in another way.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Bymonth: []int{int(time.January)},
		Dtstart: time.Date(1998, 1, 1, 9, 0, 0, 0, time.Local),
		Until:   time.Date(2000, 1, 31, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nWeekly for 10 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:    rrule.WEEKLY,
		Count:   10,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery other week, 6 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:     rrule.WEEKLY,
		Interval: 2,
		Count:    6,
		Dtstart:  time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nWeekly on Tuesday and Thursday for 5 weeks.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Count:     10,
		Wkst:      rrule.SU,
		Byweekday: []rrule.Weekday{rrule.TU, rrule.TH},
		Dtstart:   time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery other week on Tuesday and Thursday, for 8 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Interval:  2,
		Count:     8,
		Wkst:      rrule.SU,
		Byweekday: []rrule.Weekday{rrule.TU, rrule.TH},
		Dtstart:   time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nMonthly on the 1st Friday for ten occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.MONTHLY,
		Count:     10,
		Byweekday: []rrule.Weekday{rrule.FR.Nth(1)},
		Dtstart:   time.Date(1997, 9, 5, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery other month on the 1st and last Sunday of the month for 10 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.MONTHLY,
		Interval:  2,
		Count:     10,
		Byweekday: []rrule.Weekday{rrule.SU.Nth(1), rrule.SU.Nth(-1)},
		Dtstart:   time.Date(1997, 9, 7, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nMonthly on the second to last Monday of the month for 6 months.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.MONTHLY,
		Count:     6,
		Byweekday: []rrule.Weekday{rrule.MO.Nth(-2)},
		Dtstart:   time.Date(1997, 9, 22, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nMonthly on the third to the last day of the month, for 6 months.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:       rrule.MONTHLY,
		Count:      6,
		Bymonthday: []int{-3},
		Dtstart:    time.Date(1997, 9, 28, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nMonthly on the 2nd and 15th of the month for 5 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:       rrule.MONTHLY,
		Count:      5,
		Bymonthday: []int{2, 15},
		Dtstart:    time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nMonthly on the first and last day of the month for 3 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:       rrule.MONTHLY,
		Count:      5,
		Bymonthday: []int{-1, 1},
		Dtstart:    time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery 18 months on the 10th thru 15th of the month for 10 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:       rrule.MONTHLY,
		Interval:   18,
		Count:      10,
		Bymonthday: []int{10, 11, 12, 13, 14, 15},
		Dtstart:    time.Date(1997, 9, 10, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery Tuesday, every other month, 6 occurences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.MONTHLY,
		Interval:  2,
		Count:     6,
		Byweekday: []rrule.Weekday{rrule.TU},
		Dtstart:   time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nYearly in June and July for 4 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:    rrule.YEARLY,
		Count:   4,
		Bymonth: []int{6, 7},
		Dtstart: time.Date(1997, 6, 10, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery 3rd year on the 1st, 100th and 200th day for 4 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.YEARLY,
		Count:     4,
		Interval:  3,
		Byyearday: []int{1, 100, 200},
		Dtstart:   time.Date(1997, 1, 1, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery 20th Monday of the year, 3 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.YEARLY,
		Count:     3,
		Byweekday: []rrule.Weekday{rrule.MO.Nth(20)},
		Dtstart:   time.Date(1997, 5, 19, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nMonday of week number 20 (where the default start of the week is Monday), 3 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.YEARLY,
		Count:     3,
		Byweekno:  []int{20},
		Byweekday: []rrule.Weekday{rrule.MO},
		Dtstart:   time.Date(1997, 5, 12, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nThe week number 1 may be in the last year.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Count:     3,
		Byweekno:  []int{1},
		Byweekday: []rrule.Weekday{rrule.MO},
		Dtstart:   time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nAnd the week numbers greater than 51 may be in the next year.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Count:     3,
		Byweekno:  []int{52},
		Byweekday: []rrule.Weekday{rrule.SU},
		Dtstart:   time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nOnly some years have week number 53:")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Count:     3,
		Byweekno:  []int{53},
		Byweekday: []rrule.Weekday{rrule.MO},
		Dtstart:   time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery Friday the 13th, 4 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:       rrule.YEARLY,
		Count:      4,
		Byweekday:  []rrule.Weekday{rrule.FR},
		Bymonthday: []int{13},
		Dtstart:    time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery four years, the first Tuesday after a Monday in November, 3 occurrences (U.S. Presidential Election day).")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:       rrule.YEARLY,
		Interval:   4,
		Count:      3,
		Bymonth:    []int{11},
		Byweekday:  []rrule.Weekday{rrule.TU},
		Bymonthday: []int{2, 3, 4, 5, 6, 7, 8},
		Dtstart:    time.Date(1996, 11, 5, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nThe 3rd instance into the month of one of Tuesday, Wednesday or Thursday, for the next 3 months.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.MONTHLY,
		Count:     3,
		Byweekday: []rrule.Weekday{rrule.TU, rrule.WE, rrule.TH},
		Bysetpos:  []int{3},
		Dtstart:   time.Date(1997, 9, 4, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nThe 2nd to last weekday of the month, 3 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.MONTHLY,
		Count:     3,
		Byweekday: []rrule.Weekday{rrule.MO, rrule.TU, rrule.WE, rrule.TH, rrule.FR},
		Bysetpos:  []int{-2},
		Dtstart:   time.Date(1997, 9, 29, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery 3 hours from 9:00 AM to 5:00 PM on a specific day.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:     rrule.HOURLY,
		Interval: 3,
		Dtstart:  time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local),
		Until:    time.Date(1997, 9, 2, 17, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery 15 minutes for 6 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:     rrule.MINUTELY,
		Interval: 15,
		Count:    6,
		Dtstart:  time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery hour and a half for 4 occurrences.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:     rrule.MINUTELY,
		Interval: 90,
		Count:    4,
		Dtstart:  time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nEvery 20 minutes from 9:00 AM to 4:40 PM for two days.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:     rrule.MINUTELY,
		Interval: 20,
		Count:    48,
		Byhour:   []int{9, 10, 11, 12, 13, 14, 15, 16},
		Byminute: []int{0, 20, 40},
		Dtstart:  time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println("\nAn example where the days generated makes a difference because of wkst.")
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Interval:  2,
		Count:     4,
		Byweekday: []rrule.Weekday{rrule.TU, rrule.SU},
		Wkst:      rrule.MO,
		Dtstart:   time.Date(1997, 8, 5, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}

	fmt.Println()
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.WEEKLY,
		Interval:  2,
		Count:     4,
		Byweekday: []rrule.Weekday{rrule.TU, rrule.SU},
		Wkst:      rrule.SU,
		Dtstart:   time.Date(1997, 8, 5, 9, 0, 0, 0, time.Local)})
	for _, v := range r.All() {
		fmt.Println(v)
	}
	fmt.Println(r)
}

func testRRuleSet() {
	fmt.Println("\nDaily, for 7 days, jumping Saturday and Sunday occurrences.")
	set := rrule.Set{}
	r, _ := rrule.NewRRule(rrule.ROption{
		Freq:    rrule.DAILY,
		Count:   7,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	set.RRule(r)
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:      rrule.YEARLY,
		Byweekday: []rrule.Weekday{rrule.SA, rrule.SU},
		Dtstart:   time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	set.ExRule(r)
	for _, v := range set.All() {
		fmt.Println(v)
	}

	fmt.Println("\nWeekly, for 4 weeks, plus one time on day 7, and not on day 16.")
	set = rrule.Set{}
	r, _ = rrule.NewRRule(rrule.ROption{
		Freq:    rrule.WEEKLY,
		Count:   4,
		Dtstart: time.Date(1997, 9, 2, 9, 0, 0, 0, time.Local)})
	set.RRule(r)
	set.RDate(time.Date(1997, 9, 7, 9, 0, 0, 0, time.Local))
	set.ExDate(time.Date(1997, 9, 16, 9, 0, 0, 0, time.Local))
	for _, v := range set.All() {
		fmt.Println(v)
	}
}

func main() {
	testRRule()
	testRRuleSet()
}
