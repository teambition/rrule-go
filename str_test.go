package rrule

import (
	"testing"
)

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

func TestSetStr(t *testing.T) {
	setStr := "RRULE:FREQ=DAILY;UNTIL=20180517T235959Z\n" +
		"RRULE:FREQ=WEEKLY;INTERVAL=2;BYDAY=MO,TU\n" +
		"EXRULE:FREQ=WEEKLY;INTERVAL=4;BYDAY=MO\n" +
		"EXDATE;VALUE=DATE-TIME:20180525T070000Z,20180530T130000Z\n" +
		"RDATE;VALUE=DATE-TIME:20180801T131313Z,20180902T141414Z\n"

	set, err := StrToRRuleSet(setStr)
	if err != nil {
		t.Fatalf("StrToRRuleSet(%s) returned error: %v", setStr, err)
	}

	// matching parsed RRules
	rRules := set.GetRRule()
	if len(rRules) != 2 {
		t.Errorf("Unexpected number of rrule parsed: %v != 2, rrules: %v", len(rRules), rRules)
	}
	if rRules[0].String() != "FREQ=DAILY;UNTIL=20180517T235959Z" {
		t.Errorf("Unexpected rrule: %s", rRules[0].String())
	}
	if rRules[1].String() != "FREQ=WEEKLY;INTERVAL=2;BYDAY=MO,TU" {
		t.Errorf("Unexpected rrule: %s", rRules[0].String())
	}

	// matching parsed EXRules
	exRules := set.GetExRule()
	if len(exRules) != 1 {
		t.Errorf("Unexpected number of exrule parsed: %v != 1, exrules: %v", len(exRules), exRules)
	}
	if exRules[0].String() != "FREQ=WEEKLY;INTERVAL=4;BYDAY=MO" {
		t.Errorf("Unexpected exrule: %s", exRules[0].String())
	}

	// matching parsed EXDates
	exDates := set.GetExDate()
	if len(exDates) != 2 {
		t.Errorf("Unexpected number of exDates: %v != 2, %v", len(exDates), exDates)
	}
	if [2]string{timeToStr(exDates[0]), timeToStr(exDates[1])} != [2]string{"20180525T070000Z", "20180530T130000Z"} {
		t.Errorf("Unexpected exDates: %v", exDates)
	}

	// matching parsed RDates
	rDates := set.GetRDate()
	if len(rDates) != 2 {
		t.Errorf("Unexpected number of rDates: %v != 2, %v", len(rDates), rDates)
	}
	if [2]string{timeToStr(rDates[0]), timeToStr(rDates[1])} != [2]string{"20180801T131313Z", "20180902T141414Z"} {
		t.Errorf("Unexpected exDates: %v", exDates)
	}
}
