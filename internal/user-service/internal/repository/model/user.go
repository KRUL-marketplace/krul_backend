package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string       `db:"id"`
	Info      UserInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserInfo struct {
	PhoneNumber string `db:"phone_number"`
	Name        string `db:"name"`
}
