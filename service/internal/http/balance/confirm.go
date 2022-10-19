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

func (u *balance) Confirm(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.BALANCE + c.SYSTEM_ERROR + err.Error()),
		}, http.StatusInternalServerError)

		return
	}
	balanceConfirmIn, err := validateConfirm(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.BALANCE + c.SERVICE_ERROR + err.Error()),
		}, http.StatusBadRequest)
		return
	}

	_, err = u.balanceService.Confirm(balanceConfirmIn)
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

func validateConfirm(bodyBytes []byte) (*models.BalanceConfirmFields, error) {
	confirmIn := models.BalanceConfirmFields{}
	confirm := models.Balance{}

	err := json.Unmarshal(bodyBytes, &confirm)
	if err != nil {
		return nil, fmt.Errorf(c.JSON_PARSE_ERROR)
	}
	fmt.Printf("confirm: %v\n", confirm)

	processID, err := vo.ExamineIntID(confirm.ProcessID)
	if err != nil {
		return nil, fmt.Errorf(c.PROCESS_ID + err.Error())
	}
	confirmIn.ProcessID = processID

	return &confirmIn, nil
}
