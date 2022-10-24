package accounting_test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	h "user_balance/service/internal/http/accounting"
	logger "user_balance/service/internal/logger/mocks"
	"user_balance/service/internal/models"
	serviceAccounting "user_balance/service/internal/service/accounting/mocks"
	"user_balance/service/internal/vo"

	"github.com/stretchr/testify/assert"
)

type responseTest struct {
	body       models.AccountingListOut
	statusCode int
}

func (r *responseTest) Header() http.Header {
	return http.Header{}
}
func (r *responseTest) Write(in []byte) (int, error) {
	err := json.Unmarshal(in, &r.body)
	if err != nil {
		return 0, err
	}
	return 0, nil
}
func (r *responseTest) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}

func request(in string) *http.Request {
	stringReader := strings.NewReader(in)
	stringReadCloser := io.NopCloser(stringReader)
	return &http.Request{Body: stringReadCloser}
}

func Test_accounting_List(t *testing.T) {

	assert := assert.New(t)
	w := responseTest{}

	tests := []struct {
		name            string
		r               *http.Request
		expectedSuccess bool
		expectedCode    int
		expectedError   *string
	}{
		{
			name:            "validate",
			r:               request(`{"month":10,"year":2022}`),
			expectedSuccess: false,
			expectedCode:    400,
			expectedError:   nil,
		},
	}

	mockLogger := logger.NewLogger(t)
	mockAccounting := serviceAccounting.NewAccounting(t)
	input := models.AccountingListIn{}

	var m uint64 = 10
	month, _ := vo.ExamineMonth(m)
	input.Month = month
	var y uint64 = 2022
	year, _ := vo.ExamineYear(y)
	input.Year = year
	mockAccounting.On("List", input).Return(nil, nil).Once()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := h.New(mockAccounting, mockLogger)
			u.List(&w, tt.r)
			assert.Equal(tt.expectedSuccess, w.body.Success)
			assert.Equal(tt.expectedCode, w.statusCode)
			if tt.expectedError != nil {
				if w.body.Error == nil {
					t.Error("expectedError != nil, but error == nil")
					return
				}
				if *w.body.Error != *tt.expectedError {
					t.Errorf("error = %v, expected %v", *w.body.Error, *tt.expectedError)
				}
			} else {
				if tt.expectedError == nil {
					if w.body.Error != nil {
						t.Error("expectedError = nil, but error != nil")
						return
					}
				}
			}
		})

	}
}
