package application

import (
	"go-daily-log/domain/log"
	domain "go-daily-log/domain/repository"
	infr "go-daily-log/infrastructure/db/sqlite/repository"
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
	l.entryRepository.Save(logEntry)
}
