package services

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidFormatMessage = errors.New("invalid repeat format")
	layout                  = "20060102"
	timeNow                 = time.Now()
)

func nextDate(now time.Time, dstart string, repeat string) (int, error) {
	err := repeatCheck(repeat)
	if err != nil {
		return 0, err
	}

	dStartFormat, err := time.Parse(layout, dstart)
	if err != nil {
		return 0, errors.New("invalid date format")
	}

	nextDate, err := nextDateAfter(now, dStartFormat, repeat)
	if err != nil {
		return 0, err
	}

	nextDateInt, err := strconv.Atoi(nextDate.Format(layout))
	if err != nil {
		return 0, err
	}

	return nextDateInt, nil
}

func nextDateAfter(now, start time.Time, repeat string) (time.Time, error) {
	repeatSplit := strings.Fields(strings.TrimSpace(repeat))

	if len(repeatSplit) == 0 {
		return time.Time{}, errors.New("repeat value is empty")
	}

	switch repeatSplit[0] {
	case "d":
		return nextDays(now, start, repeatSplit[1])
	case "y":
		return nextYear(now, start)
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

	for {
		start = start.AddDate(0, 0, numb)
		if start.After(now) {
			break
		}
	}

	return start, nil
}

func nextYear(now, start time.Time) (time.Time, error) {
	for {
		start = start.AddDate(1, 0, 0)
		if start.After(now) {
			break
		}
	}

	return start, nil
}

func repeatCheck(repeat string) error {
	if repeat == "" {
		return errors.New("repeat value is empty")
	}

	repeatTrim := strings.TrimSpace(repeat)
	repeatSplit := strings.Split(repeatTrim, " ")
	errMessage := errors.New("invalid repeat format")

	letter := repeatSplit[0]

	if letter == "d" || letter == "y" {
		strList := strings.Split(repeatSplit[1], ",")
		for _, v := range strList {
			_, err := strconv.Atoi(v)
			if err != nil {
				return errMessage
			}
		}
		return nil
	}

	return errMessage
}

func afterNow(date, now time.Time) bool {
	return date.After(now)
}
