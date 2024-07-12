package application

import (
	"go-daily-log/domain/log"
	domain "go-daily-log/domain/repository"
	infr "go-daily-log/infrastructure/db/sqlite/repository"
	log2 "log"
)

type ExportMonthCommand struct {
	entryRepository domain.EntryRepository
}

func ExportMonth() *ExportMonthCommand {
	return &ExportMonthCommand{
		entryRepository: infr.NewEntryRepositoryImpl(),
	}
}

func (e *ExportMonthCommand) Apply(year int, month int) []log.Entry {
	log2.Println("Export monthly log to csv", year, month)
	return e.entryRepository.FindByMonth(year, month)
}
