package model

import (
	"database/sql"
	"time"
)

type Category struct {
	ID        uint32
	Info      CategoryInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type CategoryInfo struct {
	Name string
	Slug string
}
