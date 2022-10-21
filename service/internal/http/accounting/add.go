package accounting

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

func (u *accounting) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.ACCOUNTING + c.SYSTEM_ERROR + err.Error()),
		}, http.StatusInternalServerError)

		return
	}
	input, err := validateList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.ACCOUNTING + c.VALIDATE_ERROR + err.Error()),
		}, http.StatusBadRequest)
		return
	}

	out, err := u.accountingService.List(input)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.ACCOUNTING + c.SERVICE_ERROR + err.Error()),
		}, http.StatusForbidden)
		return
	}

	helpers.HttpResponse(w, models.AccountingListOut{
		Success: true,
		Accounting: out,
	}, http.StatusOK)
}

func validateList(bodyBytes []byte) (*models.AccountingListIn, error) {
	listIn := models.AccountingListIn{}
	list := models.AccountingList{}

	err := json.Unmarshal(bodyBytes, &list)
	if err != nil {
		return nil, fmt.Errorf(c.JSON_PARSE_ERROR)
	}

	month, err := vo.ExamineMonth(list.Month)
	if err != nil {
		return nil, fmt.Errorf(c.MONTH + err.Error())
	}
	listIn.Month = month

	year, err := vo.ExamineYear(list.Year)
	if err != nil {
		return nil, fmt.Errorf(c.MONTH + err.Error())
	}
	listIn.Year = year

	return &listIn, nil
}
