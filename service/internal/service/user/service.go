package user

import (
	"database/sql"
	"user_balance/service/internal/logger"
	"user_balance/service/internal/models"
	"user_balance/service/internal/repository"
)

type User interface {
	Add(user *models.UserFieldsAdd) (sql.Result, error)
}

type user struct {
	hub    repository.Hub
	logger logger.Logger
}

func New(hub repository.Hub, logger logger.Logger) User {
	return &user{hub: hub, logger: logger}
}
