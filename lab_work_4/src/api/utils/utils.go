package utils

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

func ConvertBodyStringDateToTime(s string) (time.Time, error) {

	dateParts := strings.Split(s, ".")
	if len(dateParts) != 3 {
		return time.Time{}, errors.New("invalid date format. Date should be in the format 'dd.mm.yyyy'")
	}

	day, err := strconv.Atoi(dateParts[0])
	if err != nil || day < 1 || day > 31 {
		return time.Time{}, errors.New("invalid day")
	}

	month, err := strconv.Atoi(dateParts[1])
	if err != nil || month < 1 || month > 12 {
		return time.Time{}, errors.New("invalid month")
	}

	year, err := strconv.Atoi(dateParts[2])
	if err != nil || year < 0 {
		return time.Time{}, errors.New("invalid year")
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil
}

func ConvertStringParamToIntegerID(s string) (int, error) {

	idInt, err := strconv.Atoi(s)

	if err != nil {
		log.Println("Failed to convert id to int:", err)
		return 0, errors.New("format failed: failed to convert id to int")
	}

	return idInt, nil
}
