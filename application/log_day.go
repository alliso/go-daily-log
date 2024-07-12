package application

import (
	"fmt"
	"go-daily-log/domain/log"
	domain "go-daily-log/domain/repository"
	infr "go-daily-log/infrastructure/db/sqlite/repository"
	"time"
)

type LogDayCommand struct {
	entryRepository domain.EntryRepository
}

func LogDay() *LogDayCommand {
	return &LogDayCommand{
		entryRepository: infr.NewEntryRepositoryImpl(),
	}
}

func (l *LogDayCommand) Apply(logEntry log.Entry) {
	fmt.Println(logEntry.Date.Format(time.DateOnly), logEntry.Project, logEntry.Task)
	l.entryRepository.Save(logEntry)
}
