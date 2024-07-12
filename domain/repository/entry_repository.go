package repository

import (
	"go-daily-log/domain/log"
	"time"
)

type EntryRepository interface {
	Save(entry log.Entry)

	FindByDay(date time.Time) []log.Entry

	FindByMonth(year int, month int) []log.Entry

	GetDaysInclusiveFrom(days int, date time.Time) []log.Entry
}
