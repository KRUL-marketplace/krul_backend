package app

import (
	"context"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/client/db"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/client/db/pg"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/client/db/transaction"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/api"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/config"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/repository"
	categoryRepository "github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/repository/category"
	"github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/service"
	categoryService "github.com/KRUL-marketplace/krul_backend/internal/category-service/internal/service/category"
	"log"
)

type serviceProvider struct {
	categoryRepository repository.CategoryRepository
	categoryService    service.CategoryService

	grpcConfig config.GRPCConfig
	httpConfig config.HTTPConfig
	pgConfig   config.PGConfig

	dbClient  db.Client
	txManager db.TxManager

	categoryImpl *api.Implementation
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

func (s *serviceProvider) CategoryRepository(ctx context.Context) repository.CategoryRepository {
	if s.categoryRepository == nil {
		s.categoryRepository = categoryRepository.NewRepository(s.DBClient(ctx))
	}

	return s.categoryRepository
}

func (s *serviceProvider) CategoryService(ctx context.Context) service.CategoryService {
	if s.categoryService == nil {
		s.categoryService = categoryService.NewService(
			s.CategoryRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.categoryService
}

func (s *serviceProvider) CategoryImpl(ctx context.Context) *api.Implementation {
	if s.categoryImpl == nil {
		s.categoryImpl = api.NewImplementation(s.CategoryService(ctx))
	}

	return s.categoryImpl
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
