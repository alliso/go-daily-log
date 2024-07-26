package application

import (
	domain2 "go-daily-log/domain"
	"go-daily-log/domain/log"
	domain "go-daily-log/domain/repository"
	"go-daily-log/infrastructure"
	infr "go-daily-log/infrastructure/db/sqlite/repository"
)

type ExportMonthCommand struct {
	entryRepository domain.EntryRepository
	exporter        domain2.Exporter
}

func ExportMonth() *ExportMonthCommand {
	return &ExportMonthCommand{
		entryRepository: infr.NewEntryRepositoryImpl(),
		exporter:        infrastructure.CsvCSVExporter{},
	}
}

func (e *ExportMonthCommand) Apply(year int, month int) []log.Entry {
	entries := e.entryRepository.FindByMonth(year, month)
	return e.exporter.Export(entries, year, month)
}
