package common

import "time"

type Time struct {
	time.Time
}

func ParseTimestamp(stamp int64) Time {
	return Time{time.Unix(0, stamp*(int64(time.Millisecond)/int64(time.Nanosecond)))}
}

func (t Time) Timestamp() int64 {
	return t.UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}
