package helpers

import (
	log2 "log"
	"strconv"
)

func GetYearAndMonth(strYear string, strMonth string) (int, int) {
	var err error
	var year int
	var month int
	year, err = strconv.Atoi(strYear)
	month, err = strconv.Atoi(strMonth)
	if err != nil {
		log2.Fatal("Invalid year, month")
	}

	return year, month
}
