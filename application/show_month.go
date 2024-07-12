package application

import (
	"go-daily-log/domain/log"
	domain "go-daily-log/domain/repository"
	infr "go-daily-log/infrastructure/db/sqlite/repository"
	log2 "log"
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
	log2.Println("Show monthly log ", year, month)
	return s.entryRepository.FindByMonth(year, month)
}
