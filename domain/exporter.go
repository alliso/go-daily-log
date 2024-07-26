package domain

import "go-daily-log/domain/log"

type Exporter interface {
	Export(entry []log.Entry, year int, mont int) []log.Entry
}
