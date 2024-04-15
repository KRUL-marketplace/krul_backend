package api

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/converter"
	desc "github.com/KRUL-marketplace/krul_backend/internal/product_service/pkg/product-service"
	"log"
)

func (i *Implementation) GetAll(ctx context.Context, _ *desc.GetAllRequest) (*desc.GetAllResponse, error) {
	products, err := i.productService.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var productMessages []*desc.Product

	for _, product := range products {
		log.Printf("id: %s, name: %s, description: %s, price: %d, created_at: %v, updated_at:%v\n",
			product.ID, product.Info.Name, product.Info.Description, product.Info.Description)

		productMessage := converter.ToProductFromService(product)
		productMessages = append(productMessages, productMessage)
	}

	response := &desc.GetAllResponse{
		Product: productMessages,
	}
	
	return response, nil
}
