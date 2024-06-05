package app

import (
	"context"
	"github.com/ShevelevEvgeniy/app/internal/usecase"
	"github.com/ShevelevEvgeniy/app/pkg/retry_func"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"os"

	"github.com/ShevelevEvgeniy/app/config"
	kptHandler "github.com/ShevelevEvgeniy/app/internal/http-server/api/v1/handlers/kpt_handler"
	landPlotsHandler "github.com/ShevelevEvgeniy/app/internal/http-server/api/v1/handlers/land_plots_handler"
	"github.com/ShevelevEvgeniy/app/internal/repository"
	kptRepository "github.com/ShevelevEvgeniy/app/internal/repository/kpt_repository"
	landPlotsRepository "github.com/ShevelevEvgeniy/app/internal/repository/land_plots_repository"
	s3Client "github.com/ShevelevEvgeniy/app/internal/s3_client"
	minioClient "github.com/ShevelevEvgeniy/app/internal/s3_client/minio_client"
	services "github.com/ShevelevEvgeniy/app/internal/service"
	kptService "github.com/ShevelevEvgeniy/app/internal/service/kpt_service"
	landPlotsService "github.com/ShevelevEvgeniy/app/internal/service/land_plots_service"
	"github.com/ShevelevEvgeniy/app/internal/validations"
	"github.com/ShevelevEvgeniy/app/lib/logger/sl"
	dbConnection "github.com/ShevelevEvgeniy/app/pkg/db_connection"
	"github.com/go-playground/validator/v10"
)

type DiContainer struct {
	log                 *slog.Logger
	cfg                 *config.Config
	dbConn              *pgxpool.Pool
	landPlotsRepository repository.LandPlotsRepository
	kptRepository       repository.KptRepository
	s3Clients           s3Client.MinioClient
	landPlotsService    services.LandPlotsService
	kptService          services.KptService
	validator           *validator.Validate
	retry               *retry_func.RetryFunc
	kptUseCase          kptHandler.KptUseCaseInterface
	landPlotsHandler    *landPlotsHandler.LandPlotsHandler
	kptHandler          *kptHandler.KptHandler
}

func NewDiContainer(log *slog.Logger) *DiContainer {
	return &DiContainer{
		log: log,
	}
}

func (d *DiContainer) Config(_ context.Context) *config.Config {
	if d.cfg == nil {
		d.cfg = config.MustLoad(d.log)
	}

	return d.cfg
}

func (d *DiContainer) DbConn(ctx context.Context) *pgxpool.Pool {
	if d.dbConn == nil {
		var err error
		d.dbConn, err = dbConnection.Connect(ctx, d.Config(ctx).DB)
		if err != nil {
			d.log.Error("Failed to initialize db connection: ", sl.Err(err))
			os.Exit(1)
		}
		return d.dbConn
	}

	return d.dbConn
}

func (d *DiContainer) LandPlotsRepository(ctx context.Context) repository.LandPlotsRepository {
	if d.landPlotsRepository == nil {
		d.landPlotsRepository = landPlotsRepository.NewLandPlotsRepository(d.DbConn(ctx))
	}

	return d.landPlotsRepository
}

func (d *DiContainer) KptRepository(ctx context.Context) repository.KptRepository {
	if d.kptRepository == nil {
		d.kptRepository = kptRepository.NewKptRepository(d.DbConn(ctx))
	}

	return d.kptRepository
}

func (d *DiContainer) S3Clients(ctx context.Context) s3Client.MinioClient {
	if d.s3Clients == nil {
		var err error
		d.s3Clients, err = minioClient.NewMinioClient(ctx, d.Config(ctx))
		if err != nil {
			d.log.Error("Failed to initialize clients: ", sl.Err(err))
			os.Exit(1)
		}
	}

	return d.s3Clients
}

func (d *DiContainer) LandPlotsService(ctx context.Context) services.LandPlotsService {
	if d.landPlotsService == nil {
		d.landPlotsService = landPlotsService.NewLandPlotsService(d.LandPlotsRepository(ctx))
	}

	return d.landPlotsService
}

func (d *DiContainer) KptService(ctx context.Context) services.KptService {
	if d.kptService == nil {
		d.kptService = kptService.NewKptService(d.KptRepository(ctx), d.S3Clients(ctx))
	}

	return d.kptService
}

func (d *DiContainer) Validator() *validator.Validate {
	if d.validator == nil {
		d.validator = validator.New()

		err := validations.RegisterValidations(d.validator)
		if err != nil {
			d.log.Error("Failed to register validations: ", sl.Err(err))
			os.Exit(1)
		}
	}

	return d.validator
}

func (d *DiContainer) Retry(ctx context.Context) *retry_func.RetryFunc {
	if d.retry == nil {
		d.retry = retry_func.NewRetryFunc(d.Config(ctx).RetryConfig, d.log)
	}

	return d.retry
}

func (d *DiContainer) KptUseCase(ctx context.Context) kptHandler.KptUseCaseInterface {
	if d.kptUseCase == nil {
		d.kptUseCase = usecase.NewKptUseCase(d.KptService(ctx), d.LandPlotsService(ctx), d.Retry(ctx), d.log)
	}

	return d.kptUseCase
}

func (d *DiContainer) LandPlotsHandler(ctx context.Context) *landPlotsHandler.LandPlotsHandler {
	if d.landPlotsHandler == nil {
		d.landPlotsHandler = landPlotsHandler.NewLandPlotsHandler(d.log, d.LandPlotsService(ctx), d.Validator())
	}

	return d.landPlotsHandler
}

func (d *DiContainer) KptHandler(ctx context.Context) *kptHandler.KptHandler {
	if d.kptHandler == nil {
		d.kptHandler = kptHandler.NewKptHandler(d.log, d.KptUseCase(ctx))
	}

	return d.kptHandler
}
