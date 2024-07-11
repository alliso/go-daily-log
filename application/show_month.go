package application

import (
	"go-daily-log/domain/log"
	log2 "log"
)

func ShowMonth(year int, month int) []log.Entry {
	log2.Println("Show monthly log ", year, month)
	return nil
}
