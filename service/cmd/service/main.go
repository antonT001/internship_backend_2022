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
	userService "user_balance/service/internal/service/user"

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

	userHandle := userHttp.New(userService, logger)
	balanceHandle := balanceHttp.New(balanceService, logger)

	router := mux.NewRouter()
	router.Use(commonMiddleware.Handle)

	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/add", userHandle.Add).Methods(http.MethodPost)

	balanceRouter := router.PathPrefix("/balance").Subrouter()
	balanceRouter.HandleFunc("/add", balanceHandle.Add).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("ready")
	log.Fatal(srv.ListenAndServe())
}
