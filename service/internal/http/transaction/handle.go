package transaction

import (
	"net/http"
	"user_balance/service/internal/logger"
	transactionService "user_balance/service/internal/service/transaction"
)

type Transaction interface {
	Add(w http.ResponseWriter, r *http.Request)
	Confirm(w http.ResponseWriter, r *http.Request)
	Cancel(w http.ResponseWriter, r *http.Request)
	List(w http.ResponseWriter, r *http.Request)
}

type transaction struct {
	transactionService transactionService.Transaction
	logger             logger.Logger
}

func New(transactionService transactionService.Transaction, logger logger.Logger) Transaction {
	return &transaction{transactionService: transactionService, logger: logger}
}
