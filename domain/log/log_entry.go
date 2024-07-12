package log

import "time"

type Entry struct {
	Id      int64
	Date    time.Time
	Project string
	Task    string
	Hours   float64
}
