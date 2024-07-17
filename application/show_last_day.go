package application

import (
	domain "go-daily-log/domain/repository"
	infr "go-daily-log/infrastructure/db/sqlite/repository"
	"time"
)

type GetLastDayCommand struct {
	entryRepository domain.EntryRepository
}

func LastDay() *GetLastDayCommand {
	return &GetLastDayCommand{
		entryRepository: infr.NewEntryRepositoryImpl(),
	}
}

func (s *GetLastDayCommand) Get() time.Time {
	return s.entryRepository.GetLastDay()
}
