package application

import (
	"go-daily-log/domain/log"
	domain "go-daily-log/domain/repository"
	infr "go-daily-log/infrastructure/db/sqlite/repository"
)

type ShowMonthCommand struct {
	entryRepository domain.EntryRepository
}

func ShowMonth() *ShowMonthCommand {
	return &ShowMonthCommand{
		entryRepository: infr.NewEntryRepositoryImpl(),
	}
}

func (s *ShowMonthCommand) Apply(year int, month int) []log.Entry {
	return s.entryRepository.FindByMonth(year, month)
}
