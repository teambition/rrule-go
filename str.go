package rrule

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// DateTimeFormat is date-time format used in iCalendar (RFC 5545)
	DateTimeFormat = "20060102T150405Z"
	// LocalDateTimeFormat is a date-time format without Z prefix
	LocalDateTimeFormat = "20060102T150405"
	// DateFormat is date format used in iCalendar (RFC 5545)
	DateFormat = "20060102"
)

func timeToStr(time time.Time) string {
	return time.UTC().Format(DateTimeFormat)
}

func timeToDtStartStr(time time.Time) string {
	return fmt.Sprintf("TZID=%s:%s", time.Location().String(), time.Format(LocalDateTimeFormat))
}

func strToTime(str string) (time.Time, error) {
	return strToTimeInLoc(str, time.UTC)
}

func strToTimeInLoc(str string, loc *time.Location) (time.Time, error) {
	if len(str) == len(DateFormat) {
		return time.ParseInLocation(DateFormat, str, loc)
	}
	if len(str) == len(LocalDateTimeFormat) {
		return time.ParseInLocation(LocalDateTimeFormat, str, loc)
	}
	// date-time format carries zone info
	return time.Parse(DateTimeFormat, str)
}

func (f Frequency) String() string {
	return [...]string{
		"YEARLY", "MONTHLY", "WEEKLY", "DAILY",
		"HOURLY", "MINUTELY", "SECONDLY"}[f]
}

func strToFreq(str string) (Frequency, error) {
	freqMap := map[string]Frequency{
		"YEARLY": YEARLY, "MONTHLY": MONTHLY, "WEEKLY": WEEKLY, "DAILY": DAILY,
		"HOURLY": HOURLY, "MINUTELY": MINUTELY, "SECONDLY": SECONDLY,
	}
	result, ok := freqMap[str]
	if !ok {
		return 0, errors.New("undefined frequency: " + str)
	}
	return result, nil
}

func (wday Weekday) String() string {
	s := [...]string{"MO", "TU", "WE", "TH", "FR", "SA", "SU"}[wday.weekday]
	if wday.n == 0 {
		return s
	}
	return fmt.Sprintf("%+d%s", wday.n, s)
}

func strToWeekday(str string) (Weekday, error) {
	if len(str) < 2 {
		return Weekday{}, errors.New("undefined weekday: " + str)
	}
	weekMap := map[string]Weekday{
		"MO": MO, "TU": TU, "WE": WE, "TH": TH,
		"FR": FR, "SA": SA, "SU": SU}
	result, ok := weekMap[str[len(str)-2:]]
	if !ok {
		return Weekday{}, errors.New("undefined weekday: " + str)
	}
	if len(str) > 2 {
		n, e := strconv.Atoi(str[:len(str)-2])
		if e != nil {
			return Weekday{}, e
		}
		result.n = n
	}
	return result, nil
}

func strToWeekdays(value string) ([]Weekday, error) {
	contents := strings.Split(value, ",")
	result := make([]Weekday, len(contents))
	var e error
	for i, s := range contents {
		result[i], e = strToWeekday(s)
		if e != nil {
			return nil, e
		}
	}
	return result, nil
}

func appendIntsOption(options []string, key string, value []int) []string {
	if len(value) == 0 {
		return options
	}
	valueStr := make([]string, len(value))
	for i, v := range value {
		valueStr[i] = strconv.Itoa(v)
	}
	return append(options, fmt.Sprintf("%s=%s", key, strings.Join(valueStr, ",")))
}

func strToInts(value string) ([]int, error) {
	contents := strings.Split(value, ",")
	result := make([]int, len(contents))
	var e error
	for i, s := range contents {
		result[i], e = strconv.Atoi(s)
		if e != nil {
			return nil, e
		}
	}
	return result, nil
}

func (option *ROption) String() string {
	result := []string{fmt.Sprintf("FREQ=%v", option.Freq)}
	if !option.Dtstart.IsZero() && !option.RFC {
		result = append(result, fmt.Sprintf("DTSTART=%s", timeToStr(option.Dtstart)))
	}
	if option.Interval != 0 {
		result = append(result, fmt.Sprintf("INTERVAL=%v", option.Interval))
	}
	if option.Wkst != MO {
		result = append(result, fmt.Sprintf("WKST=%v", option.Wkst))
	}
	if option.Count != 0 {
		result = append(result, fmt.Sprintf("COUNT=%v", option.Count))
	}
	if !option.Until.IsZero() {
		result = append(result, fmt.Sprintf("UNTIL=%v", timeToStr(option.Until)))
	}
	result = appendIntsOption(result, "BYSETPOS", option.Bysetpos)
	result = appendIntsOption(result, "BYMONTH", option.Bymonth)
	result = appendIntsOption(result, "BYMONTHDAY", option.Bymonthday)
	result = appendIntsOption(result, "BYYEARDAY", option.Byyearday)
	result = appendIntsOption(result, "BYWEEKNO", option.Byweekno)
	if len(option.Byweekday) != 0 {
		valueStr := make([]string, len(option.Byweekday))
		for i, wday := range option.Byweekday {
			valueStr[i] = wday.String()
		}
		result = append(result, fmt.Sprintf("BYDAY=%s", strings.Join(valueStr, ",")))
	}
	result = appendIntsOption(result, "BYHOUR", option.Byhour)
	result = appendIntsOption(result, "BYMINUTE", option.Byminute)
	result = appendIntsOption(result, "BYSECOND", option.Bysecond)
	result = appendIntsOption(result, "BYEASTER", option.Byeaster)
	return strings.Join(result, ";")
}

// StrToROption converts string to ROption
func StrToROption(rfcString string) (*ROption, error) {
	return StrToROptionInLocation(rfcString, time.UTC)
}

// StrToROptionInLocation is same as StrToROption but in case local
// time is supplied as date-time/date field (ex. UNTIL), it is parsed
// as a time in a given location (time zone)
func StrToROptionInLocation(rfcString string, loc *time.Location) (*ROption, error) {
	rfcString = strings.TrimSpace(rfcString)
	if len(rfcString) == 0 {
		return nil, errors.New("empty string")
	}
	result := ROption{}
	result.RFC = true
	for _, attr := range strings.Split(rfcString, ";") {
		keyValue := strings.Split(attr, "=")
		if len(keyValue) != 2 {
			return nil, errors.New("wrong format")
		}
		key, value := keyValue[0], keyValue[1]
		if len(value) == 0 {
			return nil, errors.New(key + " option has no value")
		}
		var e error
		switch key {
		case "FREQ":
			result.Freq, e = strToFreq(value)
		case "DTSTART":
			result.RFC = false
			result.Dtstart, e = strToTimeInLoc(value, loc)
		case "INTERVAL":
			result.Interval, e = strconv.Atoi(value)
		case "WKST":
			result.Wkst, e = strToWeekday(value)
		case "COUNT":
			result.Count, e = strconv.Atoi(value)
		case "UNTIL":
			result.Until, e = strToTimeInLoc(value, loc)
		case "BYSETPOS":
			result.Bysetpos, e = strToInts(value)
		case "BYMONTH":
			result.Bymonth, e = strToInts(value)
		case "BYMONTHDAY":
			result.Bymonthday, e = strToInts(value)
		case "BYYEARDAY":
			result.Byyearday, e = strToInts(value)
		case "BYWEEKNO":
			result.Byweekno, e = strToInts(value)
		case "BYDAY":
			result.Byweekday, e = strToWeekdays(value)
		case "BYHOUR":
			result.Byhour, e = strToInts(value)
		case "BYMINUTE":
			result.Byminute, e = strToInts(value)
		case "BYSECOND":
			result.Bysecond, e = strToInts(value)
		case "BYEASTER":
			result.Byeaster, e = strToInts(value)
		default:
			return nil, errors.New("unknown RRULE property: " + key)
		}
		if e != nil {
			return nil, e
		}
	}
	return &result, nil
}

func (r *RRule) String() string {
	return r.OrigOptions.String()
}

func (set *Set) String() string {
	res := set.Recurrence()
	return strings.Join(res, "\n")
}

// StrToRRule converts string to RRule
func StrToRRule(rfcString string) (*RRule, error) {
	option, e := StrToROption(rfcString)
	if e != nil {
		return nil, e
	}
	return NewRRule(*option)
}

// StrToRRuleSet converts string to RRuleSet
func StrToRRuleSet(s string) (*Set, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, errors.New("empty string")
	}
	ss := strings.Split(s, "\n")
	return StrSliceToRRuleSet(ss)
}

// StrSliceToRRuleSet converts given str slice to RRuleSet
func StrSliceToRRuleSet(ss []string) (*Set, error) {
	set := Set{}

	// According to RFC DTSTART is always the first line.
	firstName, err := processRRuleName(ss[0])
	if err != nil {
		return nil, err
	}

	if firstName == "DTSTART" {
		nameLen := strings.IndexAny(ss[0], ";:")
		dt, err := strToDtStart(ss[0][nameLen+1:])
		if err != nil {
			return nil, fmt.Errorf("strToDtStart failed: %v", err)
		}
		set.DTStart(dt)
		// We've processed the first one
		ss = ss[1:]
	}

	for _, line := range ss {

		name, err := processRRuleName(line)
		if err != nil {
			return nil, err
		}

		nameLen := len(name)

		switch name {
		case "RRULE", "EXRULE":
			r, err := StrToRRule(line[nameLen+1:])

			if err != nil {
				return nil, fmt.Errorf("strToRRule failed: %v", err)
			}

			if !set.GetDTStart().IsZero() {
				opt := r.OrigOptions
				opt.Dtstart = set.GetDTStart()
				r, err = NewRRule(opt)
				if err != nil {
					return nil, fmt.Errorf("could not add dtstart to rule: %v", r)
				}
			}

			if name == "RRULE" {
				set.RRule(r)
			} else {
				set.ExRule(r)
			}
		case "RDATE", "EXDATE":
			ts, err := StrToDates(line[nameLen+1:])
			if err != nil {
				return nil, fmt.Errorf("strToDates failed: %v", err)
			}
			for _, t := range ts {
				if name == "RDATE" {
					set.RDate(t)
				} else {
					set.ExDate(t)
				}
			}
		default:
			return nil, fmt.Errorf("unsupported property: %v", name)
		}
	}

	return &set, nil
}

// StrToDates is intended to parse RDATE and EXDATE properties supporting only
// VALUE=DATE-TIME (DATE and PERIOD are not supported).
// Accepts string with format: "VALUE=DATE-TIME;[TZID=...]:{time},{time},...,{time}"
// or simply "{time},{time},...{time}" and parses it to array of dates
func StrToDates(str string) (ts []time.Time, err error) {
	tmp := strings.Split(str, ":")
	if len(tmp) > 2 {
		return nil, fmt.Errorf("bad format")
	}
	var loc *time.Location
	if len(tmp) == 2 {
		params := strings.Split(tmp[0], ";")
		for _, param := range params {
			if strings.HasPrefix(param, "TZID=") {
				loc, err = parseTZID(param)
			} else if param != "VALUE=DATE-TIME" {
				err = fmt.Errorf("unsupported: %v", param)
			}
			if err != nil {
				return nil, fmt.Errorf("bad dates param: %s", err.Error())
			}
		}
		tmp = tmp[1:]
	}
	for _, datestr := range strings.Split(tmp[0], ",") {
		var t time.Time
		if loc == nil {
			t, err = strToTime(datestr)
		} else {
			t, err = strToTimeInLoc(datestr, loc)
		}
		if err != nil {
			return nil, fmt.Errorf("strToTime failed: %v", err)
		}
		ts = append(ts, t)
	}
	return
}

// processRRuleName processes the name of an RRule off a multi-line RRule set
func processRRuleName(line string) (string, error) {
	line = strings.ToUpper(strings.TrimSpace(line))
	if line == "" {
		return "", fmt.Errorf("bad format %v", line)
	}

	nameLen := strings.IndexAny(line, ";:")
	if nameLen <= 0 {
		return "", fmt.Errorf("bad format %v", line)
	}

	name := line[:nameLen]
	if strings.IndexAny(name, "=") > 0 {
		return "", fmt.Errorf("bad format %v", line)
	}

	return name, nil
}

// strToDtStart accepts string with format: "(TZID={timezone}:)?{time}" and parses it to a date
// may be used to parse DTSTART rules, without the DTSTART; part.
func strToDtStart(str string) (time.Time, error) {
	tmp := strings.Split(str, ":")
	if len(tmp) > 2 || len(tmp) == 0 {
		return time.Time{}, fmt.Errorf("bad format")
	}

	if len(tmp) == 2 {
		// tzid
		loc, err := parseTZID(tmp[0])
		if err != nil {
			return time.Time{}, err
		}
		return strToTimeInLoc(tmp[1], loc)
	}
	// no tzid, len == 1
	return strToTime(tmp[0])
}

func parseTZID(s string) (*time.Location, error) {
	if !strings.HasPrefix(s, "TZID=") || len(s) == len("TZID=") {
		return nil, fmt.Errorf("bad TZID parameter format")
	}
	return time.LoadLocation(s[len("TZID="):])
}
