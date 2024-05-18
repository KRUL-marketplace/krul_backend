package product

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"product-catalog-service/client/db"
	"product-catalog-service/internal/repository/category/model"
	productModel "product-catalog-service/internal/repository/product/model"
	"strconv"
	"strings"
)

func (r *repo) GetAll(ctx context.Context) ([]*productModel.GetProduct, error) {
	builder := sq.Select("p.id AS product_id", "p.name AS product_name", "p.slug AS product_slug",
		"p.description AS product_description", "p.price AS product_price", "p.created_at AS product_created_at",
		"p.updated_at AS product_updated_at",
		"STRING_AGG(c.id::character varying, ',') AS category_ids", "STRING_AGG(c.name, ',') AS categories",
		"STRING_AGG(c.slug, ',') AS category_slugs").
		From("products p").
		LeftJoin("product_categories pc ON p.id = pc.product_id").
		LeftJoin("categories c ON pc.category_id = c.id").
		GroupBy("p.id, p.name, p.description, p.price, p.created_at, p.updated_at").
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

	var products []*productModel.GetProduct
	for rows.Next() {
		var categoriesIds sql.NullString
		var categories sql.NullString
		var categorySlugs sql.NullString

		var product productModel.GetProduct
		err := rows.Scan(&product.ID, &product.Info.Name, &product.Info.Slug, &product.Info.Description, &product.Info.Price,
			&product.CreatedAt, &product.UpdatedAt, &categoriesIds, &categories, &categorySlugs)

		if err != nil {
			return nil, err
		}

		if categoriesIds.Valid {
			categoryIdsSlice := strings.Split(categoriesIds.String, ",")
			categoryNamesSlice := strings.Split(categories.String, ",")
			categorySlugsSlice := strings.Split(categorySlugs.String, ",")

			var categoryList []model.Category
			for i, name := range categoryNamesSlice {
				id, _ := strconv.ParseUint(categoryIdsSlice[i], 10, 32)
				category := model.Category{
					ID: uint32(id),
					Info: model.CategoryInfo{
						Name: name,
						Slug: categorySlugsSlice[i],
					},
				}
				categoryList = append(categoryList, category)
			}

			product.Info.Categories = categoryList
		}

		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
