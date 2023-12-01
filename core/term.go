package core

import (
	"time"
)

type Term struct {
	FromDate Date
	ToDate   Date
}

func NewTerm(fromDate Date, toDate Date) Term {
	return Term{
		FromDate: fromDate,
		ToDate:   toDate,
	}
}

type DateRange []Date

func (t *Term) GetDateRange() DateRange {
	from, _ := t.FromDate.AsTime()
	to, _ := t.ToDate.AsTime()
	var dateRange []Date
	for current := from; !current.After(to); current = current.Add(24 * time.Hour) {
		dateRange = append(dateRange, NewDate(current))
	}
	return dateRange
}
