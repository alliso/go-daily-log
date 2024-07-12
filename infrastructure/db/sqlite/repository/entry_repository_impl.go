package repository

import (
	"database/sql"
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
			VALUES (?,?,?,?)`

	if _, err := e.db.Exec(sql, entry.Task, entry.Date, entry.Project, entry.Hours); err != nil {
		panic(err)
	}
}

func (e *EntryRepositoryImpl) FindByDay(date time.Time) []log.Entry {

	sql := `SELECT * FROM log_entries WHERE date = ?`

	var entries []log.Entry

	if rows, err := e.db.Query(sql, date.Format(time.DateOnly)); err != nil {
		panic(err)
	} else {
		for rows.Next() {
			var entry log.Entry
			if err := rows.Scan(&entry.Date, &entry.Project, &entry.Task, &entry.Hours); err != nil {
				panic(err)
			}
			entries = append(entries, entry)
		}

	}

	return entries
}
