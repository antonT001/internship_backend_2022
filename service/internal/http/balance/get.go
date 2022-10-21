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

func (u *balance) Get(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.BALANCE + c.SYSTEM_ERROR + err.Error()),
		}, http.StatusInternalServerError)

		return
	}
	input, err := validateGet(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.BALANCE + c.VALIDATE_ERROR + err.Error()),
		}, http.StatusBadRequest)
		return
	}

	get, err := u.balanceService.Get(input)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.BALANCE + c.SERVICE_ERROR + err.Error()),
		}, http.StatusForbidden)
		return
	}

	helpers.HttpResponse(w, models.BalanceGetOut{
		Success: true,
		Balance: get,
	}, http.StatusOK)
}

func validateGet(bodyBytes []byte) (*vo.IntID, error) {
	add := models.BalanceGetIn{}
	err := json.Unmarshal(bodyBytes, &add)
	if err != nil {
		return nil, fmt.Errorf(c.JSON_PARSE_ERROR)
	}
	fmt.Printf("add: %v\n", add)

	user_id, err := vo.ExamineIntID(add.UserID)
	if err != nil {
		return nil, fmt.Errorf(c.ID + err.Error())
	}

	return &user_id, nil
}
