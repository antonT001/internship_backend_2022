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
	balanceService "user_balance/service/internal/service/balance"
	transactionService "user_balance/service/internal/service/transaction"
	userService "user_balance/service/internal/service/user"

	transactionHttp "user_balance/service/internal/http/transaction"
	balanceHttp "user_balance/service/internal/http/balance"
	userHttp "user_balance/service/internal/http/user"

	"github.com/gorilla/mux"
)

func main() {
	logger := logger.New()
	db := clients.New(logger)
	hub := repository.New(db, logger)

	commonMiddleware := middlewares.NewCommonMiddleware()

	userService := userService.New(hub, logger)
	balanceService := balanceService.New(hub, logger)
	transactionService := transactionService.New(hub, logger)

	userHandle := userHttp.New(userService, logger)
	balanceHandle := balanceHttp.New(balanceService, logger)
	transactionHandle := transactionHttp.New(transactionService, logger)

	router := mux.NewRouter()
	router.Use(commonMiddleware.Handle)

	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/add", userHandle.Add).Methods(http.MethodPost)

	balanceRouter := router.PathPrefix("/balance").Subrouter()
	balanceRouter.HandleFunc("/add", balanceHandle.Add).Methods(http.MethodPost)
	balanceRouter.HandleFunc("/get", balanceHandle.Get).Methods(http.MethodPost)

	transactionRouter := router.PathPrefix("/transaction").Subrouter()
	transactionRouter.HandleFunc("/pay", transactionHandle.Pay).Methods(http.MethodPost)
	transactionRouter.HandleFunc("/confirm", transactionHandle.Confirm).Methods(http.MethodPost)
	transactionRouter.HandleFunc("/cancel", transactionHandle.Cancel).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("ready")
	log.Fatal(srv.ListenAndServe())
}
