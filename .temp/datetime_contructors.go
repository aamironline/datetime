package temp

import "time"

// Now returns the local DateTime corresponding to the current time.
func Now() DateTime {
	return DateTime(time.Now())
}

func DateFromTS(ts int64) DateTime {
	return DateTime(time.Unix(ts, 0))
}

func DateFromTime(t time.Time) DateTime {
	return DateTime(t)
}

/**
 * The following constructors are create for convenience.
 * They include constructors for Today, Daystart and Dayend.
 */

// Return today's DateTime at 00:00:00
func DayStart(dt ...time.Time) DateTime {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return DateTime(start)
}

// Returns today's DateTime at 11:59 PM
func DayEnd(dt ...time.Time) DateTime {
	var t time.Time
	if len(dt) > 0 {
		t = dt[0]
	} else {
		t = time.Now()
	}
	end := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
	return DateTime(end)
}

// Today is an alias for DayStart
func Today(dt ...time.Time) DateTime {
	return DayStart(dt...)
}

// Yesterday returns the local DateTime corresponding to yesterday's date.
func Yesterday() DateTime {
	return DateTime(Now().AddDate(0, 0, -1))
}

// Tomorrow returns the local DateTime corresponding to tomorrow's date.
func Tomorrow() DateTime {
	return DateTime(Now().AddDate(0, 0, 1))
}

/**
 * The following constructors are from the time package, they
 * are included here for convenience.
 */

// Date returns the Time corresponding to
//
//	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
//
// in the appropriate zone for that time in the given location.
//
// The month, day, hour, min, sec, and nsec values may be outside
// their usual ranges and will be normalized during the conversion.
// For example, October 32 converts to November 1.
//
// A daylight savings time transition skips or repeats times.
// For example, in the United States, March 13, 2011 2:15am never occurred,
// while November 6, 2011 1:15am occurred twice. In such cases, the
// choice of time zone, and therefore the time, is not well-defined.
// Date returns a time that is correct in one of the two zones involved
// in the transition, but it does not guarantee which.
//
// Date panics if loc is nil.
func Date(year int, month time.Month, day, hour, min, sec, nsec int, loc *time.Location) DateTime {
	return DateTime(time.Date(year, month, day, hour, min, sec, nsec, loc))
}

// Unix returns the local DateTime corresponding to the given Unix time,
// sec seconds and nsec nanoseconds since January 1, 1970 UTC.
// It is valid to pass nsec outside the range [0, 999999999].
// Not all sec values have a corresponding time value. One such
// value is 1<<63-1 (the largest int64 value).
func Unix(sec int64, nsec int64) DateTime {
	return DateTime(time.Unix(sec, nsec))
}

// UnixMilli returns the local DateTime corresponding to the given Unix time,
// msec milliseconds since January 1, 1970 UTC.
func UnixMilli(msec int64) DateTime {
	return DateTime(time.UnixMilli(msec))
}

// UnixMicro returns the local DateTime corresponding to the given Unix time,
// usec microseconds since January 1, 1970 UTC.
func UnixMicro(usec int64) DateTime {
	return DateTime(time.UnixMicro(usec))
}
