package application

import (
	"go-daily-log/domain/log"
	domain "go-daily-log/domain/repository"
	infra "go-daily-log/infrastructure/db/sqlite/repository"
	log2 "log"
	"time"
)

type ShowWeekCommand struct {
	entryRepository domain.EntryRepository
}

func ShowWeek() *ShowWeekCommand {
	return &ShowWeekCommand{
		entryRepository: infra.NewEntryRepositoryImpl(),
	}
}

func (s *ShowWeekCommand) Apply(date time.Time) []log.Entry {
	log2.Println("Show weekly log ", date)

	for wd := date.Weekday(); wd != time.Monday; {
		date = date.AddDate(0, 0, -1)
		wd = date.Weekday()
	}

	return s.entryRepository.GetDaysInclusiveFrom(6, date)
}
