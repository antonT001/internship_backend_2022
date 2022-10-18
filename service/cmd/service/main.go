package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"user_balance/service/internal/clients"
	"user_balance/service/internal/logger"
	"user_balance/service/internal/repository"
	userService "user_balance/service/internal/service/user"

	userHttp "user_balance/service/internal/http/user"

	"github.com/gorilla/mux"
)

func main() {
	logger := logger.New()
	db := clients.New(logger)
	hub := repository.New(db, logger)

	userService := userService.New(hub, logger)

	userHandle := userHttp.New(userService, logger)

	router := mux.NewRouter()
	userRouter := router.PathPrefix("/user").Subrouter()

	userRouter.HandleFunc("/add", userHandle.Add).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("ready")
	log.Fatal(srv.ListenAndServe())
}
