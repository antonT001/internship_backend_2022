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

func (u *balance) Add(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.BALANCE + c.SYSTEM_ERROR + err.Error()),
		}, http.StatusInternalServerError)

		return
	}
	BalanceIn, err := validateAdd(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.BALANCE + c.SERVICE_ERROR + err.Error()),
		}, http.StatusBadRequest)
		return
	}

	_, err = u.balanceService.Add(BalanceIn)
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

func validateAdd(bodyBytes []byte) (*models.BalanceFields, error) {
	addIn := models.BalanceFields{}
	add := models.Balance{}

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

	processID, err := vo.ExamineIntID(add.ProcessID)
	if err != nil {
		return nil, fmt.Errorf(c.PROCESS_ID + err.Error())
	}
	addIn.ProcessID = processID

	return &addIn, nil
}