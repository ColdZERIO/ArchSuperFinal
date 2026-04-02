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
		return nextDays(now, start, repeatSplit...)
	case "y":
		return nextYear(now, start, repeatSplit...)
	// case "w":
	// 	return time.Time{}, errors.New("invalid repeat format")
	// 	// add func for week
	// case "m":
	// 	return time.Time{}, errors.New("invalid repeat format")
	// 	// add func for month
	default:
		return time.Time{}, errors.New("invalid repeat format")
	}
}

func nextDays(now, start time.Time, parts ...string) (time.Time, error) {
	if len(parts) != 2 {
		return time.Time{}, ErrInvalidFormatMessage
	}

	numb, err := strconv.Atoi(parts[1])
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

func nextYear(now, start time.Time, parts ...string) (time.Time, error) {
	if len(parts) != 1 {
		return time.Time{}, ErrInvalidFormatMessage
	}

	newYear := now.AddDate(1, 0, 0)

	return newYear, nil
}

// func NextWeekDays(now, start time.Time, parts ...string) (time.Time, error) {
// 	if len(parts) != 2 {
// 		return time.Time{}, ErrInvalidFormatMessage
// 	}

// 	maxValue := 7

// }

func repeatCheck(repeat string) error {
	if repeat == "" {
		return errors.New("repeat value is empty")
	}

	repeatSplit := strings.Split(repeat, " ")
	errMessage := errors.New("invelid repeat format")

	letter := repeatSplit[0]

	if letter != "" && letter != "d" || letter != "y" { // || letter != "w" || letter != "m"
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
