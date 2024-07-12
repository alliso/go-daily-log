package application

import (
	"go-daily-log/domain/log"
	"time"
)

func ShowDay(date time.Time) []log.Entry {
	return []log.Entry{
		{
			Date:    time.Now(),
			Project: "delivery-point",
			Task:    "Soportar a Mike",
			Hours:   200,
		},
		{
			Date:    time.Now(),
			Project: "delivery-point",
			Task:    "Soportar a Mike todavia mas",
			Hours:   400,
		},
	}
}
