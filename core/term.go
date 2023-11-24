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

type Days []time.Time

func (t Term) DayDuration() Days {
	from, to := t.fromTo()
	diff := to.Sub(from)
	hours := int(diff.Hours())
	dayCount := hours / 24
	var days []time.Time
	days = append(days, from)
	for i := 1; i < dayCount; i++ {
		days = append(days, from.AddDate(0, 0, i))
	}
	days = append(days, to)
	return days
}

func (t Term) fromTo() (from time.Time, to time.Time) {
	fromDate, _ := t.FromDate.Validate()
	toDate, _ := t.ToDate.Validate()
	return *fromDate, *toDate
}
