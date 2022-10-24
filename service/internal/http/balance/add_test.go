package balance_test

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	h "user_balance/service/internal/http/balance"
	logger "user_balance/service/internal/logger/mocks"
	"user_balance/service/internal/models"
	serviceBalance "user_balance/service/internal/service/balance/mocks"
	"user_balance/service/internal/vo"

	"github.com/stretchr/testify/assert"
)

type responseTest struct {
	body       models.Out
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
			name: "validate",
			r: request(`{
				"money": 10,
				"order_id": 12345,
				"service_id": 5,
				"service_name": "test",
				"type": 0,
				"user_id": 10}
			`),
			expectedSuccess: true,
			expectedCode:    200,
			expectedError:   nil,
		},
	}

	mockLogger := logger.NewLogger(t)
	mockBalance := serviceBalance.NewBalance(t)
	input := models.TransactionFields{}

	userId, _ := vo.ExamineIntID(10)
	input.UserID = userId
	input.Type = 0 //TODO добавить vo, Type: 0 - списание, 1 - пополнение
	money, _ := vo.ExamineDeltaMoney(10)
	input.Money = money
	serviceID, _ := vo.ExamineIntID(5)
	input.ServiceID = serviceID
	serviceMame, _ := vo.ExamineName("test")
	input.ServiceName = *serviceMame
	orderID, _ := vo.ExamineIntID(12345)
	input.OrderID = orderID

	mockBalance.On("Add", input).Return(nil, nil).Once()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := h.New(mockBalance, mockLogger)
			u.Add(&w, tt.r)
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
