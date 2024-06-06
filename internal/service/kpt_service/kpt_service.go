package kpt_service

import (
	"bytes"
	"context"
	"time"

	"github.com/ShevelevEvgeniy/app/internal/dto"
	"github.com/ShevelevEvgeniy/app/internal/repository"
	"github.com/ShevelevEvgeniy/app/internal/repository/model"
	s3Client "github.com/ShevelevEvgeniy/app/internal/s3_client"
	customErrors "github.com/ShevelevEvgeniy/app/lib/custom_errors"
	"github.com/pkg/errors"
)

type KptService struct {
	repository repository.KptRepository
	Clients    s3Client.MinioClient
}

func NewKptService(repository repository.KptRepository, clients s3Client.MinioClient) *KptService {
	return &KptService{
		repository: repository,
		Clients:    clients,
	}
}

func (k *KptService) ExistKpt(ctx context.Context, kptInfo *model.Kpt) error {
	dateFormation, err := k.repository.GetKptDateFormation(ctx, kptInfo.CadQuarter)
	if err != nil {
		return errors.Wrap(err, "failed to get kpt date formation")
	}

	if dateFormation != (time.Time{}) && (dateFormation.After(kptInfo.DateFormation) || dateFormation.Equal(kptInfo.DateFormation)) {
		return customErrors.ErrKptAlreadyExist
	}

	return nil
}

func (k *KptService) SaveKptInfo(ctx context.Context, kptInfo *model.Kpt) error {
	return k.repository.SaveKptInfo(ctx, *kptInfo)
}

func (k *KptService) UploadFileToMinio(ctx context.Context, dto *dto.KptDto) error {
	return k.Clients.UploadFile(ctx, dto.KptHeaders.Filename, bytes.NewReader(dto.Kpt.Bytes()), dto.KptHeaders.Size)
}

func (k *KptService) GetKptInfo(ctx context.Context, cadQuarter string) (dto.KptInfo, error) {
	kptInfo, err := k.repository.GetKptInfo(ctx, cadQuarter)
	if err != nil {
		return dto.KptInfo{}, errors.Wrap(err, "failed to get kpt info")
	}
	kptInfo.CadQuarter = cadQuarter

	return kptInfo, nil
}

func (k *KptService) GetKptLink(ctx context.Context, cadQuarter string) (string, error) {
	kptName, err := k.repository.GetKptName(ctx, cadQuarter)
	if err != nil {
		return "", errors.Wrap(err, "failed to get kpt name")
	}

	link, err := k.Clients.GetLinkDownload(ctx, kptName)
	if err != nil {
		return "", err
	}

	return link.String(), err
}
