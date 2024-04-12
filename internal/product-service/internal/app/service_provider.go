package app

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/config"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/api"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository"
	productRepository "github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/service"
	productService "github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/service/product"
	"log"
)

type serviceProvider struct {
	productRepository repository.ProductRepository
	productService    service.ProductService

	grpcConfig  config.GRPCConfig
	productImpl *api.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) ProductRepository() repository.ProductRepository {
	if s.productRepository == nil {
		s.productRepository = productRepository.NewRepository()
	}

	return s.productRepository
}

func (s *serviceProvider) ProductService() service.ProductService {
	if s.productService == nil {
		s.productService = productService.NewService(
			s.ProductRepository(),
		)
	}

	return s.productService
}

func (s *serviceProvider) ProductImpl(ctx context.Context) *api.Implementation {
	if s.productImpl == nil {
		s.productImpl = api.NewImplementation(s.ProductService())
	}

	return s.productImpl
}
