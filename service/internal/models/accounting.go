package models

import "user_balance/service/internal/vo"

type AccountingListFields struct {
	Service_Name string `json:"service_name"`
	// Money:
	// * The format of money without kopecks is used.
	// * Example 12050=120 rubles 50 kopecks
	Money int `json:"money"`
}

type AccountingList struct {
	// Month:
	// * 1 - 12
	Month uint64 `json:"month" example:"10"`
	// Year:
	// * 2007 - this year
	Year uint64 `json:"year" example:"2022"`
}

type AccountingListIn struct {
	Month vo.Month
	Year  vo.Year
}

type AccountingListOut struct {
	Success    bool    `json:"success"`
	Accounting *string `json:"path,omitempty" example:"http://localhost:9000/static/2022/10/report102022.csv"`
	Error      *string `json:"error,omitempty"`
}
