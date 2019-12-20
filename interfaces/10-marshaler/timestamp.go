package main

import (
	"strconv"
	"time"
)

type timestamp struct {
	time.Time
}

func (ts timestamp) String() string {
	if ts.IsZero() {
		return "unknown"
	}
	const layout = "2006/01"
	return ts.Format(layout)
}

func (ts timestamp) MarshalJSON() (data []byte, _ error) {
	// ts -> integer -> ts.Unix()
	// data <- integer -> strconv.AppendInt(data, integer, 10)
	return strconv.AppendInt(data, ts.Unix(), 10), nil
}

func (ts *timestamp) UnmarshalJSON(data []byte) error {
	*ts = toTimestamp(string(data))
	return nil
}
