package dao

import (
	"moon-street/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepo interface {
	GetById(int64) (model.User, error)
	Save(model.User) (int64, error)
}
