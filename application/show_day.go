package application

import (
	"go-daily-log/domain/log"
	log2 "log"
	"time"
)

func ShowDay(date time.Time) []log.Entry {
	log2.Println("Show log day ", date.Format(time.DateOnly))
	return nil
}
