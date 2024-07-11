package application

import (
	"go-daily-log/domain/log"
	log2 "log"
	"time"
)

func ExportMonth(year int, month time.Month) []log.Entry {
	log2.Println("Export monthly log to csv", year, month)
	return nil
}
