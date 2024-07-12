package helpers

import (
	"log"
	"time"
)

func ParseDate(strDate string) time.Time {
	date, err := time.Parse(time.DateOnly, strDate)
	if err != nil {
		log.Fatal("Unable to parse date")
	}

	return date
}
