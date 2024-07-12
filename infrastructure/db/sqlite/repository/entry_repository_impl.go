package repository

import (
	"database/sql"
	"fmt"
	"go-daily-log/domain/log"
	"go-daily-log/infrastructure/db/sqlite/conf"
	"time"
)

type EntryRepositoryImpl struct {
	db *sql.DB
}

func NewEntryRepositoryImpl() *EntryRepositoryImpl {
	return &EntryRepositoryImpl{db: conf.Get()}
}

func (e *EntryRepositoryImpl) Save(entry log.Entry) {

	sql := `INSERT INTO log_entries(task, date, project, hours) 
			VALUES (?,DATE(?),?,?)`

	if _, err := e.db.Exec(sql, entry.Task, entry.Date, entry.Project, entry.Hours); err != nil {
		panic(err)
	}
}

func (e *EntryRepositoryImpl) FindByDay(date time.Time) []log.Entry {

	sql := `SELECT id, date, project, task, hours FROM log_entries WHERE date = DATE(?)`

	if rows, err := e.db.Query(sql, date.Format(time.DateOnly)); err != nil {
		panic(err)
	} else {
		return rowsToEntries(rows)
	}
}

func (e *EntryRepositoryImpl) FindByMonth(year int, month int) []log.Entry {

	sql := `SELECT id, date, project, task, hours FROM log_entries WHERE DATE(date) BETWEEN DATE(?) AND DATE(?)`

	firstday := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC).Format(time.DateOnly)
	lastday := time.Date(year, time.Month(month)+1, 0, 0, 0, 0, 0, time.UTC).Format(time.DateOnly)

	fmt.Printf("Querying from day %s to day %s\n", firstday, lastday)

	if rows, err := e.db.Query(sql, firstday, lastday); err != nil {
		panic(err)
	} else {
		return rowsToEntries(rows)
	}
}

func rowsToEntries(rows *sql.Rows) []log.Entry {
	var entries []log.Entry

	for rows.Next() {
		var entry log.Entry
		if err := rows.Scan(&entry.Id, &entry.Date, &entry.Project, &entry.Task, &entry.Hours); err != nil {
			panic(err)
		}
		entries = append(entries, entry)
	}

	return entries
}
