package application

import (
	"go-daily-log/domain/log"
	domain "go-daily-log/domain/repository"
	infr "go-daily-log/infrastructure/db/sqlite/repository"
	"time"
)

type ShowDayCommand struct {
	entryRepository domain.EntryRepository
}

func ShowDay() *ShowDayCommand {
	return &ShowDayCommand{
		entryRepository: infr.NewEntryRepositoryImpl(),
	}
}

func (s *ShowDayCommand) Apply(date time.Time) []log.Entry {
	return s.entryRepository.FindByDay(date)
}
