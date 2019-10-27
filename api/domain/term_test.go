package domain_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tkitsunai/edinet-go/api/domain"
	v1 "github.com/tkitsunai/edinet-go/api/edinet/api/v1"
	"testing"
)

func Test開始期間から終了期間までの日付のリストを取得する(t *testing.T) {
	term := domain.Term{
		FromDate: v1.FileDate("2019-08-01"),
		ToDate:   v1.FileDate("2019-08-10"),
	}

	days := term.DayDuration()
	for i := 1; i < 10; i++ {
		assert.Equal(t, fmt.Sprintf("2019-08-%02d 00:00:00 +0900 Asia/Tokyo", i), days[i-1].String())
	}
}
