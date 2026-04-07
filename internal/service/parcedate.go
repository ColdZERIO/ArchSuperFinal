package services

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidFormatMessage = errors.New("invalid repeat format")
	layout = "20060102"
	timeNow = time.Now()
)

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	err := repeatCheck(repeat)
	if err != nil {
		return "", err
	}

	dStartFormat, err := time.Parse(layout, dstart)
	if err != nil {
		return "", errors.New("invalid date format")
	}

	nextDate, err := NextDateAfter(now, dStartFormat, repeat)
	if err != nil {
		return "", err
	}

	return nextDate.Format(layout), nil
}

func NextDateAfter(now, start time.Time, repeat string) (time.Time, error) {
	repeatSplit := strings.Fields(strings.TrimSpace(repeat))

	if len(repeatSplit) == 0 {
		return time.Time{}, errors.New("repeat value is empty")
	}

	switch repeatSplit[0] {
	case "d":
		return nextDays(now, start, repeatSplit[1])
	case "y":
		return nextYear(now, start, repeatSplit[1])
	default:
		return time.Time{}, errors.New("invalid repeat format")
	}
}

func nextDays(now, start time.Time, part string) (time.Time, error) {
	numb, err := strconv.Atoi(part)
	if err != nil {
		return time.Time{}, err
	}

	if numb < 1 || numb > 400 {
		return time.Time{}, ErrInvalidFormatMessage
	}

	if start.After(now) {
		return start, nil
	}

	newDate := now.AddDate(0, 0, numb)

	return newDate, nil
}

func nextYear(now, start time.Time, part string) (time.Time, error) {

	newYear := now.AddDate(1, 0, 0)

	return newYear, nil
}

func repeatCheck(repeat string) error {
	if repeat == "" {
		return errors.New("repeat value is empty")
	}

	repeatSplit := strings.Split(repeat, " ")
	errMessage := errors.New("invelid repeat format")

	letter := repeatSplit[0]

	if letter != "" && letter != "d" || letter != "y" {
		return errMessage
	}

	strList := strings.Split(repeatSplit[1], ",")
	for _, v := range strList {
		_, err := strconv.Atoi(v)
		if err != nil {
			return errMessage
		}
	}

	return nil
}
