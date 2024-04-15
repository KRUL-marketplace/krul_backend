package product

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/client/db"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/converter"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product/model"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableName = "products"

	idColumn          = "id"
	nameColumn        = "name"
	descriptionColumn = "description"
	priceColumn       = "price"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.ProductRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, product *model.ProductInfo) (string, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, descriptionColumn, priceColumn).
		Values(product.Name, product.Description, product.Price).
		Suffix("RETURNING " + idColumn)

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "product_repository.Created",
		QueryRaw: query,
	}

	var id string
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id string) (*model.Product, error) {
	builder := sq.Select(idColumn, nameColumn, descriptionColumn, priceColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "product_repository.Get",
		QueryRaw: query,
	}

	var product model.Product
	err = r.db.DB().QueryRowContext(ctx, q, args...).
		Scan(&product.ID, &product.Info.Name, &product.Info.Description, &product.Info.Price,
			&product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToProductFromRepo(&product), nil
}
