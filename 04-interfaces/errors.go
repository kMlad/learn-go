package interfaces

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hours   int
	minutes int
	seconds int
}

func parseTime(s string) (Time, error) {
	if len(s) < 7 || len(s) > 8 {
		return Time{}, errors.New("invalid string length")
	}

	timeComponents := strings.Split(s, ":")
	if len(timeComponents) != 3 {
		return Time{}, errors.New(
			"invalid format. Please provide the string in the following format: HH:mm:ss")
	}

	var time Time

	hours, err := strconv.ParseInt(timeComponents[0], 10, 32)
	if err != nil {
		return Time{}, errors.New("cannot convert hours to a valid integer")
	}
	if hours < 0 || hours >= 24 {
		return Time{}, errors.New("hours value must be between 0 and 23")
	}

	minutes, err := strconv.ParseInt(timeComponents[1], 10, 32)
	if err != nil {
		return Time{}, errors.New("cannot convert minutes to a valid integer")
	}
	if minutes < 0 || minutes >= 60 {
		return Time{}, errors.New("minutes value must be between 0 and 59")
	}

	seconds, err := strconv.ParseInt(timeComponents[2], 10, 32) // Changed index from 1 to 2
	if err != nil {
		return Time{}, errors.New("cannot convert seconds to a valid integer")
	}
	if seconds < 0 || seconds >= 60 {
		return Time{}, errors.New("seconds value must be between 0 and 59")
	}

	time = Time{hours: int(hours), minutes: int(minutes), seconds: int(seconds)}

	return time, nil
}

func Errors() {

	// currentTime := "23:14:24"
	currentTime := "01:23:09"
	// currentTime := "12:14:AA"
	// currentTime := "25:82:24"
	// currentTime := "12:82:24"

	time, err := parseTime(currentTime)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("%v hours, %v minutes, %v seconds", time.hours, time.minutes, time.seconds)
	}

}
