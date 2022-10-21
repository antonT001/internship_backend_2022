package balance

import (
	"net/http"
	"user_balance/service/internal/logger"
	balanceService "user_balance/service/internal/service/balance"
)

type Balance interface {
	Add(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type balance struct {
	balanceService balanceService.Balance
	logger         logger.Logger
}

func New(balanceService balanceService.Balance, logger logger.Logger) Balance {
	return &balance{balanceService: balanceService, logger: logger}
}
