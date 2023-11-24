package core

import (
	"errors"
	"github.com/Songmu/go-httpdate"
	"time"
)

type FileDate string

const (
	EdinetDateFormat = "2006-01-02"
)

func NewFileDate(value time.Time) FileDate {
	return FileDate(value.Format(EdinetDateFormat))
}

func (f FileDate) String() string {
	return string(f)
}

func (f FileDate) Validate() (*time.Time, error) {
	if len(f) == 0 {
		return nil, errors.New("RequiredParameter")
	}

	jstTime, err := httpdate.Str2Time(f.String(), time.FixedZone("Asia/Tokyo", 9*60*60))

	if err != nil {
		return nil, errors.New("date format error")
	}

	return &jstTime, nil
}

func (f FileDate) Format() string {
	if formatted, err := f.Validate(); err != nil {
		return string([]rune(time.Now().Format(time.RFC3339))[:10])
	} else {
		return string([]rune(formatted.Format(time.RFC3339))[:10])
	}
}
