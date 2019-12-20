package main

import "time"

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
