package usecase_test

import (
	"github.com/samber/do"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tkitsunai/edinet-go/core"
	"github.com/tkitsunai/edinet-go/edinet"
	"github.com/tkitsunai/edinet-go/port"
	"github.com/tkitsunai/edinet-go/usecase"
	"testing"
)

var in *do.Injector

func init() {
	in = do.New()
}

type MockPort struct {
	mock.Mock
}

func (m *MockPort) Get(date core.Date) (edinet.EdinetDocumentResponse, error) {
	called := m.Called(date)
	response := called.Get(0).(edinet.EdinetDocumentResponse)
	return response, called.Error(1)
}

func (m *MockPort) GetByStore(date core.Date) (edinet.EdinetDocumentResponse, error) {
	called := m.Called(date)
	response := called.Get(0).(edinet.EdinetDocumentResponse)
	return response, called.Error(1)
}

func helper(dayCount int) []edinet.EdinetDocumentResponse {
	result := make([]edinet.EdinetDocumentResponse, dayCount)
	for i := 0; i < dayCount-1; i++ {
		result[i] = edinet.EdinetDocumentResponse{}
	}
	return result
}

func TestOverview_FindOverviewByTerm(t *testing.T) {
	mockOverview := &MockPort{}

	mockOverview.On("GetByStore", mock.Anything).Return(createEdinetResponse(), nil)
	mockOverview.On("Get", mock.Anything).Return(createEdinetResponse(), nil)

	do.ProvideValue[port.Overview](in, mockOverview)
	do.ProvideValue[port.Company](in, nil)
	target, _ := usecase.NewOverview(in)

	cases := map[string]struct {
		Term    core.Term
		Refresh bool
		Expects []edinet.EdinetDocumentResponse
	}{
		"one_day": {
			Term:    core.Term{core.Date("2023-12-01"), core.Date("2023-12-01")},
			Expects: helper(1),
		},
		"two_day": {
			Term:    core.Term{core.Date("2023-12-01"), core.Date("2023-12-02")},
			Expects: helper(2),
		},
		"three_day": {
			Term:    core.Term{core.Date("2023-12-01"), core.Date("2023-12-03")},
			Expects: helper(3),
		},
		"refresh": {
			Term:    core.Term{core.Date("2023-12-01"), core.Date("2023-12-01")},
			Refresh: true,
			Expects: helper(1),
		},
	}

	for testName, tt := range cases {
		t.Run(testName, func(t *testing.T) {
			result, err := target.FindOverviewByTerm(tt.Term, false)
			assert.Nil(t, err)
			assert.EqualValues(t, tt.Expects, result)
		})
	}
}

func createEdinetResponse() edinet.EdinetDocumentResponse {
	return edinet.EdinetDocumentResponse{
		Metadata: edinet.Metadata{},
		Results:  nil,
	}
}
