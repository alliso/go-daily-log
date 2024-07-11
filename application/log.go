package application

import (
	"fmt"
	"go-daily-log/domain/log"
	"time"
)

func LogDay(logEntry log.Entry) {
	fmt.Println(logEntry.Date.Format(time.DateOnly), logEntry.Project, logEntry.Task)
}
