package models

import "user_balance/service/internal/vo"

type AccountingListFields struct {
	Service_Name string `json:"service_name"`
	Money        int    `json:"money"`
}

type AccountingList struct {
	Month uint64 `json:"month"`
	Year  uint64 `json:"year"`
}

type AccountingListIn struct {
	Month vo.Month
	Year  vo.Year
}

type AccountingListOut struct {
	Success    bool    `json:"success"`
	Accounting *string `json:"path,omitempty"`
	Error      *string `json:"error,omitempty"`
}
