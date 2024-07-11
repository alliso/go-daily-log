package log

import "time"

type Entry struct {
	Date    time.Time
	Project string
	Task    string
}
