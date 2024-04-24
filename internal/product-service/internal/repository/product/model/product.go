package model

import (
	"database/sql"
	"time"
)

type Product struct {
	ID        string       `db:"id"`
	Info      ProductInfo  `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type ProductInfo struct {
	Name        string `db:"name"`
	Slug        string `db:"slug"`
	Description string `db:"description"`
	Price       uint32 `db:"price"`
}
