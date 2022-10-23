package balance

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

// @Summary Add
// @Tags balance
// @Description Add user balance
// @Accept json
// @Produce json
// @Param input body models.Transaction true "payload"
// @Success 200 {object} models.Out
// @Failure 400 {object} models.Out
// @Failure 403 {object} models.Out
// @Failure 500 {object} models.Out
// @Router /balance/add [post]
func (u *balance) Add(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.BALANCE + c.SYSTEM_ERROR + err.Error()),
		}, http.StatusInternalServerError)

		return
	}
	input, err := validateAdd(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.BALANCE + c.VALIDATE_ERROR + err.Error()),
		}, http.StatusBadRequest)
		return
	}

	_, err = u.balanceService.Add(input)
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

func validateAdd(bodyBytes []byte) (*models.TransactionFields, error) {
	addIn := models.TransactionFields{}
	add := models.Transaction{}

	err := json.Unmarshal(bodyBytes, &add)
	if err != nil {
		return nil, fmt.Errorf(c.JSON_PARSE_ERROR)
	}
	fmt.Printf("add: %v\n", add)

	userID, err := vo.ExamineIntID(add.UserID)
	if err != nil {
		return nil, fmt.Errorf(c.ID + err.Error())
	}
	addIn.UserID = userID

	addIn.Type = add.Type //TODO добавить vo

	money, err := vo.ExamineDeltaMoney(add.Money)
	if err != nil {
		return nil, fmt.Errorf(c.MONEY + err.Error())
	}
	addIn.Money = money

	serviceID, err := vo.ExamineIntID(add.ServiceID)
	if err != nil {
		return nil, fmt.Errorf(c.SERVICE_ID + err.Error())
	}
	addIn.ServiceID = serviceID

	serviceMame, err := vo.ExamineName(add.ServiceName)
	if err != nil {
		return nil, fmt.Errorf(c.SERVICE_NAME + err.Error())
	}
	addIn.ServiceName = *serviceMame

	orderID, err := vo.ExamineIntID(add.OrderID)
	if err != nil {
		return nil, fmt.Errorf(c.ORDER_ID + err.Error())
	}
	addIn.OrderID = orderID

	return &addIn, nil
}
