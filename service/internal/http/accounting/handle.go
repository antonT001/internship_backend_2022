package accounting

import (
	"net/http"
	"user_balance/service/internal/logger"
	accountingService "user_balance/service/internal/service/accounting"
)

type Accounting interface {
	List(w http.ResponseWriter, r *http.Request)
}

type accounting struct {
	accountingService accountingService.Accounting
	logger            logger.Logger
}

func New(accountingService accountingService.Accounting, logger logger.Logger) Accounting {
	return &accounting{accountingService: accountingService, logger: logger}
}
