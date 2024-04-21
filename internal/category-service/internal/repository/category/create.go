package category

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/client/db"
	sq "github.com/Masterminds/squirrel"

	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/repository/category/model"
)

const (
	tableName       = "category"
	idColumn        = "id"
	nameColumn      = "name"
	slugColumn      = "slug"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

func (r *repo) Create(ctx context.Context, info *model.CategoryInfo) (uint32, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, slugColumn).
		Values(info.Name, info.Slug).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "category_repository.Create",
		QueryRaw: query,
	}

	var id uint32
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
