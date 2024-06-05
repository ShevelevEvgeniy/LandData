package usecase

import (
	"context"
	errs "errors"
	"github.com/ShevelevEvgeniy/app/internal/converter"
	"github.com/ShevelevEvgeniy/app/internal/dto"
	"github.com/ShevelevEvgeniy/app/internal/service"
	customErrors "github.com/ShevelevEvgeniy/app/lib/custom_errors"
	"github.com/ShevelevEvgeniy/app/lib/logger/sl"
	retryFunc "github.com/ShevelevEvgeniy/app/pkg/retry_func"
	"golang.org/x/sync/errgroup"
	"log/slog"
)

type KptUseCase struct {
	kptService service.KptService
	lpService  service.LandPlotsService
	retry      *retryFunc.RetryFunc
	log        *slog.Logger
}

func NewKptUseCase(kptService service.KptService, lpService service.LandPlotsService, retry *retryFunc.RetryFunc, log *slog.Logger) *KptUseCase {
	return &KptUseCase{
		kptService: kptService,
		lpService:  lpService,
		retry:      retry,
		log:        log,
	}
}

func (k *KptUseCase) SaveKpt(ctx context.Context, dto *dto.KptDto) error {
	kptInfo, err := converter.ToKptInfoFromKpt(dto)
	if err != nil {
		return err
	}

	err = k.kptService.ExistKpt(ctx, kptInfo)
	if err != nil {
		if errs.Is(err, customErrors.ErrKptAlreadyExist) {
			return customErrors.ErrKptAlreadyExist
		}
		k.log.Error("Failed to exist kpt", sl.Err(err))
		return err
	}

	group, gCtx := errgroup.WithContext(ctx)

	tasks := []func(ctx context.Context) error{
		func(ctx context.Context) error { return k.kptService.SaveKptInfo(ctx, kptInfo) },
		func(ctx context.Context) error { return k.kptService.UploadFileToMinio(ctx, dto) },
		func(ctx context.Context) error { return k.lpService.SaveLandPlotsFromKpt(ctx, dto) },
	}

	for _, task := range tasks {
		iterTask := task
		group.Go(func() error {
			return k.retry.Do(gCtx, func() error {
				select {
				case <-gCtx.Done():
					return gCtx.Err()
				default:
					return iterTask(gCtx)
				}
			})
		})
	}

	if err = group.Wait(); err != nil {
		k.log.Error("Failed to save kpt", sl.Err(err))
		return err
	}

	return nil
}
