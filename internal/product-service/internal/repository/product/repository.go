package product

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/client/db"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/converter"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product/model"
	sq "github.com/Masterminds/squirrel"
	"time"
)

const (
	tableName = "products"

	idColumn          = "id"
	nameColumn        = "name"
	slugColumn        = "slug"
	descriptionColumn = "description"
	priceColumn       = "price"
	createdAtColumn   = "created_at"
	updatedAtColumn   = "updated_at"
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
		Columns(nameColumn, slugColumn, descriptionColumn, priceColumn).
		Values(product.Name, product.Slug, product.Description, product.Price).
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

func (r *repo) GetById(ctx context.Context, id string) (*model.Product, error) {
	builder := sq.Select(idColumn, nameColumn, slugColumn, descriptionColumn, priceColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "product_repository.GetById",
		QueryRaw: query,
	}

	var product model.Product
	err = r.db.DB().QueryRowContext(ctx, q, args...).
		Scan(&product.ID, &product.Info.Name, &product.Info.Slug, &product.Info.Description, &product.Info.Price,
			&product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToProductFromRepo(&product), nil
}

func (r *repo) GetBySlug(ctx context.Context, slug string) (*model.Product, error) {
	builder := sq.Select(idColumn, nameColumn, slugColumn, descriptionColumn, priceColumn, createdAtColumn, updatedAtColumn).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{slugColumn: slug}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "product_repository.GetBySlug",
		QueryRaw: query,
	}

	var product model.Product
	err = r.db.DB().QueryRowContext(ctx, q, args...).
		Scan(&product.ID, &product.Info.Name, &product.Info.Slug, &product.Info.Description, &product.Info.Price,
			&product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return converter.ToProductFromRepo(&product), nil
}

func (r *repo) GetAll(ctx context.Context) ([]*model.Product, error) {
	builder := sq.Select(idColumn, nameColumn, slugColumn, descriptionColumn, priceColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "product_repository.GetAll",
		QueryRaw: query,
	}

	rows, err := r.db.DB().QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Info.Name, &product.Info.Slug, &product.Info.Description,
			&product.Info.Price, &product.CreatedAt, &product.UpdatedAt)

		if err != nil {
			return nil, err
		}
		products = append(products, converter.ToProductFromRepo(&product))
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repo) Update(ctx context.Context, id string, info *model.ProductInfo) (string, error) {
	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, info.Name).
		Set(slugColumn, info.Slug).
		Set(descriptionColumn, info.Description).
		Set(priceColumn, info.Price).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return "", err
	}

	q := db.Query{
		Name:     "product_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return "", err
	}

	return "", nil
}

func (r *repo) DeleteById(ctx context.Context, id string) error {
	builder := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	q := db.Query{
		Name:     "product_repository.DeleteById",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil
}
