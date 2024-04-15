package app

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/client/db"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/client/db/pg"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/client/db/transaction"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/api"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/config"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository"
	productRepository "github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/repository/product"
	"github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/service"
	productService "github.com/KRUL-marketplace/krul_backend/internal/product_service/internal/service/product"
	"log"
)

type serviceProvider struct {
	productRepository repository.ProductRepository
	productService    service.ProductService

	grpcConfig config.GRPCConfig
	httpConfig config.HTTPConfig
	pgConfig   config.PGConfig

	dbClient  db.Client
	txManager db.TxManager

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

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) ProductRepository(ctx context.Context) repository.ProductRepository {
	if s.productRepository == nil {
		s.productRepository = productRepository.NewRepository(s.DBClient(ctx))
	}

	return s.productRepository
}

func (s *serviceProvider) ProductService(ctx context.Context) service.ProductService {
	if s.productService == nil {
		s.productService = productService.NewService(
			s.ProductRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.productService
}

func (s *serviceProvider) ProductImpl(ctx context.Context) *api.Implementation {
	if s.productImpl == nil {
		s.productImpl = api.NewImplementation(s.ProductService(ctx))
	}

	return s.productImpl
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}
