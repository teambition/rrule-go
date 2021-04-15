package rrule

import (
	"sort"
	"time"
)

type Next func() (value time.Time, ok bool)

type Iterator interface {
	// Next is a generator of time.Time.
	// It returns false of Ok if there is no value to generate.
	Next() (value time.Time, ok bool)

	// Previous is a generator of time.Time.
	// It returns false of Ok if there is no value to generate.
	Previous() (value time.Time, ok bool)
}

// rIterator is a iterator of RRule
type rIterator struct {
	year     int
	month    time.Month
	day      int
	hour     int
	minute   int
	second   int
	weekday  int
	ii       iterInfo
	timeset  []time.Time
	total    int
	count    int
	remain   []time.Time
	finished bool
}

func (it *rIterator) generate() {
	r := it.ii.rrule
	for len(it.remain) == 0 {
		// Get dayset with the right frequency
		dayset, start, end := it.ii.getdayset(r.freq, it.year, it.month, it.day)

		// Do the "hard" work ;-)
		filtered := false
		for _, i := range dayset[start:end] {
			if len(r.bymonth) != 0 && !contains(r.bymonth, it.ii.mmask[*i]) ||
				len(r.byweekno) != 0 && it.ii.wnomask[*i] == 0 ||
				len(r.byweekday) != 0 && !contains(r.byweekday, it.ii.wdaymask[*i]) ||
				len(it.ii.nwdaymask) != 0 && it.ii.nwdaymask[*i] == 0 ||
				len(r.byeaster) != 0 && it.ii.eastermask[*i] == 0 ||
				(len(r.bymonthday) != 0 || len(r.bynmonthday) != 0) &&
					!contains(r.bymonthday, it.ii.mdaymask[*i]) &&
					!contains(r.bynmonthday, it.ii.nmdaymask[*i]) ||
				len(r.byyearday) != 0 &&
					(*i < it.ii.yearlen &&
						!contains(r.byyearday, *i+1) &&
						!contains(r.byyearday, -it.ii.yearlen+*i) ||
						*i >= it.ii.yearlen &&
							!contains(r.byyearday, *i+1-it.ii.yearlen) &&
							!contains(r.byyearday, -it.ii.nextyearlen+*i-it.ii.yearlen)) {
				dayset[*i] = nil
				filtered = true
			}
		}
		// Output results
		if len(r.bysetpos) != 0 && len(it.timeset) != 0 {
			poslist := []time.Time{}
			for _, pos := range r.bysetpos {
				var daypos, timepos int
				if pos < 0 {
					daypos, timepos = divmod(pos, len(it.timeset))
				} else {
					daypos, timepos = divmod(pos-1, len(it.timeset))
				}
				temp := []int{}
				for _, x := range dayset[start:end] {
					if x != nil {
						temp = append(temp, *x)
					}
				}
				i, err := pySubscript(temp, daypos)
				if err != nil {
					continue
				}
				timeTemp := it.timeset[timepos]
				date := it.ii.firstyday.AddDate(0, 0, i)
				res := time.Date(date.Year(), date.Month(), date.Day(),
					timeTemp.Hour(), timeTemp.Minute(), timeTemp.Second(),
					timeTemp.Nanosecond(), timeTemp.Location())
				if !timeContains(poslist, res) {
					poslist = append(poslist, res)
				}
			}
			sort.Sort(timeSlice(poslist))
			for _, res := range poslist {
				if !r.until.IsZero() && res.After(r.until) {
					r.len = it.total
					it.finished = true
					return
				} else if !res.Before(r.dtstart) {
					it.total++
					it.remain = append(it.remain, res)
					if it.count != 0 {
						it.count--
						if it.count == 0 {
							r.len = it.total
							it.finished = true
							return
						}
					}
				}
			}
		} else {
			for _, i := range dayset[start:end] {
				if i == nil {
					continue
				}
				date := it.ii.firstyday.AddDate(0, 0, *i)
				for _, timeTemp := range it.timeset {
					res := time.Date(date.Year(), date.Month(), date.Day(),
						timeTemp.Hour(), timeTemp.Minute(), timeTemp.Second(),
						timeTemp.Nanosecond(), timeTemp.Location())
					if !r.until.IsZero() && res.After(r.until) {
						r.len = it.total
						it.finished = true
						return
					} else if !res.Before(r.dtstart) {
						it.total++
						it.remain = append(it.remain, res)
						if it.count != 0 {
							it.count--
							if it.count == 0 {
								r.len = it.total
								it.finished = true
								return
							}
						}
					}
				}
			}
		}
		// Handle frequency and interval
		fixday := false
		if r.freq == YEARLY {
			it.year += r.interval
			if it.year > MAXYEAR {
				r.len = it.total
				it.finished = true
				return
			}
			it.ii.rebuild(it.year, it.month)
		} else if r.freq == MONTHLY {
			it.month += time.Month(r.interval)
			if it.month > 12 {
				div, mod := divmod(int(it.month), 12)
				it.month = time.Month(mod)
				it.year += div
				if it.month == 0 {
					it.month = 12
					it.year--
				}
				if it.year > MAXYEAR {
					r.len = it.total
					it.finished = true
					return
				}
			}
			it.ii.rebuild(it.year, it.month)
		} else if r.freq == WEEKLY {
			if r.wkst > it.weekday {
				it.day += -(it.weekday + 1 + (6 - r.wkst)) + r.interval*7
			} else {
				it.day += -(it.weekday - r.wkst) + r.interval*7
			}
			it.weekday = r.wkst
			fixday = true
		} else if r.freq == DAILY {
			it.day += r.interval
			fixday = true
		} else if r.freq == HOURLY {
			if filtered {
				// Jump to one iteration before next day
				it.hour += ((23 - it.hour) / r.interval) * r.interval
			}
			for {
				it.hour += r.interval
				div, mod := divmod(it.hour, 24)
				if div != 0 {
					it.hour = mod
					it.day += div
					fixday = true
				}
				if len(r.byhour) == 0 || contains(r.byhour, it.hour) {
					break
				}
			}
			it.timeset = it.ii.gettimeset(r.freq, it.hour, it.minute, it.second)
		} else if r.freq == MINUTELY {
			if filtered {
				// Jump to one iteration before next day
				it.minute += ((1439 - (it.hour*60 + it.minute)) / r.interval) * r.interval
			}
			for {
				it.minute += r.interval
				div, mod := divmod(it.minute, 60)
				if div != 0 {
					it.minute = mod
					it.hour += div
					div, mod = divmod(it.hour, 24)
					if div != 0 {
						it.hour = mod
						it.day += div
						fixday = true
						filtered = false
					}
				}
				if (len(r.byhour) == 0 || contains(r.byhour, it.hour)) &&
					(len(r.byminute) == 0 || contains(r.byminute, it.minute)) {
					break
				}
			}
			it.timeset = it.ii.gettimeset(r.freq, it.hour, it.minute, it.second)
		} else if r.freq == SECONDLY {
			if filtered {
				// Jump to one iteration before next day
				it.second += (((86399 - (it.hour*3600 + it.minute*60 + it.second)) / r.interval) * r.interval)
			}
			for {
				it.second += r.interval
				div, mod := divmod(it.second, 60)
				if div != 0 {
					it.second = mod
					it.minute += div
					div, mod = divmod(it.minute, 60)
					if div != 0 {
						it.minute = mod
						it.hour += div
						div, mod = divmod(it.hour, 24)
						if div != 0 {
							it.hour = mod
							it.day += div
							fixday = true
						}
					}
				}
				if (len(r.byhour) == 0 || contains(r.byhour, it.hour)) &&
					(len(r.byminute) == 0 || contains(r.byminute, it.minute)) &&
					(len(r.bysecond) == 0 || contains(r.bysecond, it.second)) {
					break
				}
			}
			it.timeset = it.ii.gettimeset(r.freq, it.hour, it.minute, it.second)
		}
		if fixday && it.day > 28 {
			daysinmonth := daysIn(it.month, it.year)
			if it.day > daysinmonth {
				for it.day > daysinmonth {
					it.day -= daysinmonth
					it.month++
					if it.month == 13 {
						it.month = 1
						it.year++
						if it.year > MAXYEAR {
							r.len = it.total
							it.finished = true
							return
						}
					}
					daysinmonth = daysIn(it.month, it.year)
				}
				it.ii.rebuild(it.year, it.month)
			}
		}
	}
}

// next returns next occurrence and true if it exists, else zero value and false
func (it *rIterator) next() (time.Time, bool) {
	if !it.finished {
		it.generate()
	}
	if len(it.remain) == 0 {
		return time.Time{}, false
	}
	value := it.remain[0]
	it.remain = it.remain[1:]
	return value, true
}

// previous returns previous occurrence and true if it exists, else zero value and false
func (it *rIterator) previous() (time.Time, bool) {
	if !it.finished {
		it.generate()
	}
	if len(it.remain) == 0 {
		return time.Time{}, false
	}
	value := it.remain[0]
	it.remain = it.remain[1:]
	return value, true
}

func (it *rIterator) Next() (value time.Time, ok bool) {
	return it.next()
}

func (it *rIterator) Previous() (value time.Time, ok bool) {
	return it.previous()
}
