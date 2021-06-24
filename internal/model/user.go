package model

import (
	"time"
)

type User struct {
	Id        int64
	Name      string
	Status    string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
