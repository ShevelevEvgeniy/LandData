package kpt_usecase

import (
	"context"
	Dto "github.com/ShevelevEvgeniy/app/internal/dto"
	"github.com/ShevelevEvgeniy/app/internal/service"
	"github.com/ShevelevEvgeniy/app/lib/logger/sl"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"log/slog"
)

type GetKptLinkAndInfoUseCase struct {
	service service.KptService
}

func NewGetKptLinkAndInfoUseCase(service service.KptService) *GetKptLinkAndInfoUseCase {
	return &GetKptLinkAndInfoUseCase{
		service: service,
	}
}

func (h *GetKptLinkAndInfoUseCase) GetLinkAndInfoKpt(ctx context.Context, log *slog.Logger, cadQuarter string) (Dto.KptInfo, error) {
	group, gCtx := errgroup.WithContext(ctx)
	var err error
	var kptInfo Dto.KptInfo
	linkChan := make(chan string, 1)

	group.Go(func() error {
		kptInfo, err = h.service.GetKptInfo(gCtx, cadQuarter)
		if err != nil {
			log.Error("Failed to get kpt info", sl.Err(err))
			return err
		}

		return nil
	})

	group.Go(func() error {
		link, err := h.service.GetKptLink(gCtx, cadQuarter)
		if err != nil {
			log.Error("Failed to get kpt link", sl.Err(err))
			return err
		}

		linkChan <- link
		return nil
	})

	if err = group.Wait(); err != nil {
		log.Error("Failed to get kpt info and link", sl.Err(err))
		return Dto.KptInfo{}, errors.Wrap(err, "Failed to get kpt info and link")
	}

	kptInfo.Link = <-linkChan

	return kptInfo, nil
}
