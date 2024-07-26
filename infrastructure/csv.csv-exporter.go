package infrastructure

import (
	"encoding/csv"
	"go-daily-log/domain/log"
	log2 "log"
	"os"
	"strconv"
)

type CsvCSVExporter struct{}

func (c CsvCSVExporter) Export(entries []log.Entry, year int, month int) []log.Entry {
	home, err := os.UserHomeDir()
	if err != nil {
		log2.Fatal(err)
	}

	csvFile, err := os.Create(home + "/tasks_" + strconv.Itoa(year) + "_" + strconv.Itoa(month) + ".csv")
	if err != nil {
		log2.Fatal(err)
	}
	writer := csv.NewWriter(csvFile)
	writer.Comma = ';'
	data := [][]string{{"Task", "Project", "Hours", "Date"}}
	for _, entry := range entries {
		data = append(data, []string{
			entry.Task,
			entry.Project,
			strconv.Itoa(int(entry.Hours)),
			entry.Date.Format("2006-01-02"),
		})
	}

	writer.WriteAll(data)
	if err := writer.Error(); err != nil {
		log2.Panic(err)
	}

	return entries
}
