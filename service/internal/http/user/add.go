package user

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

func (u *user) Add(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.USER + c.SYSTEM_ERROR + err.Error()),
		}, http.StatusInternalServerError)

		return
	}
	userAddIn, err := validateAdd(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.USER + c.SERVICE_ERROR + err.Error()),
		}, http.StatusBadRequest)
		return
	}

	_, err = u.userService.Add(userAddIn)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.USER + c.SERVICE_ERROR + err.Error()),
		}, http.StatusForbidden)
		return
	}

	helpers.HttpResponse(w, models.Out{
		Success: true,
	}, http.StatusOK)
}

func validateAdd(bodyBytes []byte) (*models.UserFieldsAdd, error) {
	productAddIn := models.UserFieldsAdd{}
	productAdd := models.UserAdd{}

	err := json.Unmarshal(bodyBytes, &productAdd)
	if err != nil {
		return nil, fmt.Errorf(c.JSON_PARSE_ERROR)
	}

	userName, err := vo.ExamineName(productAdd.UserName)
	if err != nil {
		return nil, fmt.Errorf(c.BALANCE + err.Error())
	}
	productAddIn.UserName = *userName

	return &productAddIn, nil
}
