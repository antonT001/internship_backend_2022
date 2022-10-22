package accounting

import (
	"reflect"
	"strconv"
	"user_balance/service/internal/models"
)

func (c *accounting) List(input *models.AccountingListIn) (*string, error) {
	list, err := c.hub.Accounting().List(input)
	if err != nil {
		return nil, err
	}

	records := make([][]string, len(list))
	for i, v := range list {
		records[i] = make([]string, 0, reflect.ValueOf(v).NumField())
		records[i] = append(records[i], v.Service_Name)
		records[i] = append(records[i], strconv.Itoa(v.Money))
	}

	yearStr := strconv.Itoa(int(input.Year.Year()))
	monthStr := strconv.Itoa(int(input.Month.Month()))
	path := "/" + yearStr + "/" + monthStr + "/" + "report" + monthStr + yearStr + ".csv"

	err = c.objectStorage.SaveCsvFile(path, records)
	if err != nil {
		return nil, err
	}

	return &path, nil
}
