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

// @Summary Confirm
// @Tags transaction
// @Description Transaction confirmation, debiting money in favor of the company
// @Accept json
// @Produce json
// @Param input body models.TransactionConfirm true "payload"
// @Success 200 {object} models.Out
// @Failure 400 {object} models.Out
// @Failure 403 {object} models.Out
// @Failure 500 {object} models.Out
// @Router /transaction/confirm [post]
func (u *transaction) Confirm(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.TRANSACTIONS + c.SYSTEM_ERROR + err.Error()),
		}, http.StatusInternalServerError)

		return
	}
	input, err := validateConfirm(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.TRANSACTIONS + c.VALIDATE_ERROR + err.Error()),
		}, http.StatusBadRequest)
		return
	}

	_, err = u.transactionService.Confirm(input)
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

func validateConfirm(bodyBytes []byte) (*models.TransactionConfirmFields, error) {
	confirmIn := models.TransactionConfirmFields{}
	confirm := models.TransactionConfirm{}

	err := json.Unmarshal(bodyBytes, &confirm)
	if err != nil {
		return nil, fmt.Errorf(c.JSON_PARSE_ERROR)
	}
	fmt.Printf("confirm: %v\n", confirm)

	userID, err := vo.ExamineIntID(confirm.UserID)
	if err != nil {
		return nil, fmt.Errorf(c.ID + err.Error())
	}
	confirmIn.UserID = userID

	money, err := vo.ExamineDeltaMoney(confirm.Money)
	if err != nil {
		return nil, fmt.Errorf(c.MONEY + err.Error())
	}
	confirmIn.Money = money

	serviceID, err := vo.ExamineIntID(confirm.ServiceID)
	if err != nil {
		return nil, fmt.Errorf(c.SERVICE_ID + err.Error())
	}
	confirmIn.ServiceID = serviceID

	serviceName, err := vo.ExamineName(confirm.ServiceName)
	if err != nil {
		return nil, fmt.Errorf(c.SERVICE_NAME + err.Error())
	}
	confirmIn.ServiceName = *serviceName

	orderID, err := vo.ExamineIntID(confirm.OrderID)
	if err != nil {
		return nil, fmt.Errorf(c.ORDER_ID + err.Error())
	}
	confirmIn.OrderID = orderID

	return &confirmIn, nil
}
