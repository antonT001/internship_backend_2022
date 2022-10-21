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

func (u *transaction) List(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.TRANSACTIONS + c.SYSTEM_ERROR + err.Error()),
		}, http.StatusInternalServerError)

		return
	}
	input, err := validateList(bodyBytes)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.TRANSACTIONS + c.VALIDATE_ERROR + err.Error()),
		}, http.StatusBadRequest)
		return
	}

	out, err := u.transactionService.List(input)
	if err != nil {
		helpers.HttpResponse(w, models.Out{
			Success: false,
			Error:   helpers.StringPointer(c.TRANSACTIONS + c.SERVICE_ERROR + err.Error()),
		}, http.StatusForbidden)
		return
	}

	helpers.HttpResponse(w, models.TransactionListOut{
		Success:     true,
		Transaction: out,
	}, http.StatusOK)
}

func validateList(bodyBytes []byte) (*models.TransactionListIn, error) {
	listIn := models.TransactionListIn{}
	list := models.TransactionList{}

	err := json.Unmarshal(bodyBytes, &list)
	if err != nil {
		return nil, fmt.Errorf(c.JSON_PARSE_ERROR)
	}
	fmt.Printf("list: %v\n", list)

	pageNum, err := vo.ExaminePointerIntID(list.PageNum)
	if err != nil {
		return nil, fmt.Errorf(c.ID + err.Error())
	}
	listIn.PageNum = pageNum

	userID, err := vo.ExamineIntID(list.UserID)
	if err != nil {
		return nil, fmt.Errorf(c.ID + err.Error())
	}
	listIn.UserID = userID

	if list.Filter == nil {
		return &listIn, nil
	}

	filter := models.TransactionFilterIn{}

	orderBy, err := vo.ExamineOrderBy(list.Filter.OrderBy, vo.OrderByTransactionList)
	if err != nil {
		return nil, fmt.Errorf(c.ORDER_BY + err.Error())
	}
	filter.OrderBy = orderBy

	orderDirection, err := vo.ExamineOrderDirection(list.Filter.OrderDirection)
	if err != nil {
		return nil, fmt.Errorf(c.ORDER_DIRECTION + err.Error())
	}
	filter.OrderDirection = orderDirection

	listIn.Filter = &filter

	return &listIn, nil
}
