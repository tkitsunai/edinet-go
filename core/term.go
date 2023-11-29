package core

import (
	"time"
)

type Term struct {
	FromDate FileDate
	ToDate   FileDate
}

func NewTerm(fromDate FileDate, toDate FileDate) Term {
	return Term{
		FromDate: fromDate,
		ToDate:   toDate,
	}
}

type DateRange []time.Time

func (t *Term) GetDateRange() DateRange {
	from, to := t.fromTo()
	var dateRange []time.Time
	for current := from; !current.After(to); current = current.Add(24 * time.Hour) {
		dateRange = append(dateRange, current)
	}
	return dateRange
}

func (t *Term) fromTo() (from time.Time, to time.Time) {
	fromDate, _ := t.FromDate.Validate()
	toDate, _ := t.ToDate.Validate()
	return *fromDate, *toDate
}
