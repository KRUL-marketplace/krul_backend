package model

import (
	"database/sql"
	"time"
)

type Product struct {
	ID        string
	Info      ProductInfo
	CreatedAt time.Time
	UpdateAt  sql.NullTime
}

type ProductInfo struct {
	Name        string
	Description string
	Price       uint32
}
