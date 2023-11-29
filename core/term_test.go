package core_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tkitsunai/edinet-go/core"
	"testing"
)

func Test開始期間から終了期間までが三日間の場合三つの日付リストが取得できる(t *testing.T) {
	term := core.Term{
		FromDate: core.FileDate("2019-08-01"),
		ToDate:   core.FileDate("2019-08-03"),
	}

	dateRange := term.GetDateRange()

	for i, date := range dateRange {
		assert.Equal(t, fmt.Sprintf("2019-08-%02d 00:00:00 +0900 Asia/Tokyo", i+1), date.String())
	}
}

func Test開始期間から終了期間までが二日間の場合二つの日付リストが取得できる(t *testing.T) {
	term := core.Term{
		FromDate: core.FileDate("2023-11-29"),
		ToDate:   core.FileDate("2023-11-30"),
	}

	dateRange := term.GetDateRange()

	assert.Equal(t, 2, len(dateRange))
	assert.Equal(t, "2023-11-29 00:00:00 +0900 Asia/Tokyo", dateRange[0].String())
	assert.Equal(t, "2023-11-30 00:00:00 +0900 Asia/Tokyo", dateRange[1].String())
}

func Test開始と終了期間が同じ場合は一つの日付リストが取得できる(t *testing.T) {
	term := core.Term{
		FromDate: core.FileDate("2023-11-29"),
		ToDate:   core.FileDate("2023-11-29"),
	}

	dateRange := term.GetDateRange()

	assert.Equal(t, 1, len(dateRange))
	expected := fmt.Sprintf("2023-11-29 00:00:00 +0900 Asia/Tokyo")
	for _, date := range dateRange {
		assert.Equal(t, expected, date.String())
	}
}
