package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"user_balance/service/internal/clients"
	"user_balance/service/internal/logger"
	"user_balance/service/internal/middlewares"
	"user_balance/service/internal/repository"
	accountingService "user_balance/service/internal/service/accounting"
	balanceService "user_balance/service/internal/service/balance"
	transactionService "user_balance/service/internal/service/transaction"

	accountingHttp "user_balance/service/internal/http/accounting"
	balanceHttp "user_balance/service/internal/http/balance"
	transactionHttp "user_balance/service/internal/http/transaction"

	"github.com/gorilla/mux"
)

func main() {
	logger := logger.New()
	db := clients.New(logger)
	hub := repository.New(db, logger)

	commonMiddleware := middlewares.NewCommonMiddleware()

	accountingService := accountingService.New(hub, logger)
	balanceService := balanceService.New(hub, logger)
	transactionService := transactionService.New(hub, logger)

	accountingHandle := accountingHttp.New(accountingService, logger)
	balanceHandle := balanceHttp.New(balanceService, logger)
	transactionHandle := transactionHttp.New(transactionService, logger)

	router := mux.NewRouter()
	router.Use(commonMiddleware.Handle)

	accountingRouter := router.PathPrefix("/accounting").Subrouter()
	accountingRouter.HandleFunc("/list", accountingHandle.List).Methods(http.MethodPost)

	balanceRouter := router.PathPrefix("/balance").Subrouter()
	balanceRouter.HandleFunc("/add", balanceHandle.Add).Methods(http.MethodPost)
	balanceRouter.HandleFunc("/get", balanceHandle.Get).Methods(http.MethodPost)

	transactionRouter := router.PathPrefix("/transaction").Subrouter()
	transactionRouter.HandleFunc("/pay", transactionHandle.Pay).Methods(http.MethodPost)
	transactionRouter.HandleFunc("/confirm", transactionHandle.Confirm).Methods(http.MethodPost)
	transactionRouter.HandleFunc("/cancel", transactionHandle.Cancel).Methods(http.MethodPost)
	transactionRouter.HandleFunc("/list", transactionHandle.List).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("ready")
	log.Fatal(srv.ListenAndServe())
}
