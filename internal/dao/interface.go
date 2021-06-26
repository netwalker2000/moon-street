package dao

import (
	"moon-street/internal/model"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepo interface {
	GetByName(string) (model.User, error)
	Save(model.User) (int64, error)
}
