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

func (u *transaction) Cancel(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.TRANSACTIONS + c.SYSTEM_ERROR + err.Error()),
		}, http.StatusInternalServerError)

		return
	}
	input, err := validateCancel(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.TRANSACTIONS + c.VALIDATE_ERROR + err.Error()),
		}, http.StatusBadRequest)
		return
	}

	_, err = u.transactionService.Cancel(input)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.TRANSACTIONS + c.SERVICE_ERROR + err.Error()),
		}, http.StatusForbidden)
		return
	}

	helpers.HttpResponse(w, models.Out{
		Success: true,
	}, http.StatusOK)
}

func validateCancel(bodyBytes []byte) (*models.TransactionConfirmFields, error) {
	cancelIn := models.TransactionConfirmFields{}
	cancel := models.TransactionConfirm{}

	err := json.Unmarshal(bodyBytes, &cancel)
	if err != nil {
		return nil, fmt.Errorf(c.JSON_PARSE_ERROR)
	}
	fmt.Printf("cancel: %v\n", cancel)

	userID, err := vo.ExamineIntID(cancel.UserID)
	if err != nil {
		return nil, fmt.Errorf(c.ID + err.Error())
	}
	cancelIn.UserID = userID

	money, err := vo.ExamineDeltaMoney(cancel.Money)
	if err != nil {
		return nil, fmt.Errorf(c.MONEY + err.Error())
	}
	cancelIn.Money = money

	serviceID, err := vo.ExamineIntID(cancel.ServiceID)
	if err != nil {
		return nil, fmt.Errorf(c.SERVICE_ID + err.Error())
	}
	cancelIn.ServiceID = serviceID

	orderID, err := vo.ExamineIntID(cancel.OrderID)
	if err != nil {
		return nil, fmt.Errorf(c.ORDER_ID + err.Error())
	}
	cancelIn.OrderID = orderID

	return &cancelIn, nil
}
