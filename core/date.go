package core

import (
	"time"
)

type Date string

const (
	ISO8601Basic = "2006-01-02"
)

func NewDate(value time.Time) Date {
	return Date(value.Format(ISO8601Basic))
}

func (f Date) String() string {
	return string(f)
}

func (f Date) AsTime() (time.Time, error) {
	parseTime, err := time.Parse(ISO8601Basic, f.String())
	if err != nil {
		return time.Now(), err
	}
	return parseTime, nil
}
