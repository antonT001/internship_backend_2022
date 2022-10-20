package transaction

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	c "user_balance/service/internal/constants"
	"user_balance/service/internal/helpers"
	"user_balance/service/internal/models"
	"user_balance/service/internal/vo"
)

func (u *transaction) Pay(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.BALANCE + c.SYSTEM_ERROR + err.Error()),
		}, http.StatusInternalServerError)

		return
	}
	transactionPayIn, err := validatePay(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.BALANCE + c.VALIDATE_ERROR + err.Error()),
		}, http.StatusBadRequest)
		return
	}

	_, err = u.transactionService.Pay(transactionPayIn)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.BALANCE + c.SERVICE_ERROR + err.Error()),
		}, http.StatusForbidden)
		return
	}

	helpers.HttpResponse(w, models.Out{
		Success: true,
	}, http.StatusOK)
}

func validatePay(bodyBytes []byte) (*models.TransactionFields, error) {
	payIn := models.TransactionFields{}
	pay := models.Transaction{}

	err := json.Unmarshal(bodyBytes, &pay)
	if err != nil {
		return nil, fmt.Errorf(c.JSON_PARSE_ERROR)
	}
	fmt.Printf("pay: %v\n", pay)

	userID, err := vo.ExamineIntID(pay.UserID)
	if err != nil {
		return nil, fmt.Errorf(c.ID + err.Error())
	}
	payIn.UserID = userID

	payIn.Type = pay.Type //TODO добавить vo

	money, err := vo.ExamineDeltaMoney(pay.Money)
	if err != nil {
		return nil, fmt.Errorf(c.MONEY + err.Error())
	}
	payIn.Money = money

	serviceID, err := vo.ExamineIntID(pay.ServiceID)
	if err != nil {
		return nil, fmt.Errorf(c.SERVICE_ID + err.Error())
	}
	payIn.ServiceID = serviceID

	serviceMame, err := vo.ExamineName(pay.ServiceName)
	if err != nil {
		return nil, fmt.Errorf(c.SERVICE_NAME + err.Error())
	}
	payIn.ServiceName = *serviceMame

	orderID, err := vo.ExamineIntID(pay.OrderID)
	if err != nil {
		return nil, fmt.Errorf(c.ORDER_ID + err.Error())
	}
	payIn.OrderID = orderID

	return &payIn, nil
}
