package save_kpt_handler

import (
	"bytes"
	"context"
	"encoding/xml"
	errs "errors"
	"log/slog"
	"mime/multipart"
	"net/http"

	Dto "github.com/ShevelevEvgeniy/app/internal/dto"
	"github.com/ShevelevEvgeniy/app/lib/api/response"
	customErrors "github.com/ShevelevEvgeniy/app/lib/custom_errors"
	"github.com/ShevelevEvgeniy/app/lib/logger/sl"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/pkg/errors"
)

type KptUseCaseInterface interface {
	SaveKpt(ctx context.Context, dto *Dto.KptDto) error
}

type SaveKptHandler struct {
	log     *slog.Logger
	UseCase KptUseCaseInterface
}

func NewKptHandler(log *slog.Logger, useCase KptUseCaseInterface) *SaveKptHandler {
	return &SaveKptHandler{
		log:     log,
		UseCase: useCase,
	}
}

func (k *SaveKptHandler) SaveKpt(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "KptHandler.SaveKpt"

		k.log = k.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		k.log.Info("Received HTTP POST request", slog.String("path", r.URL.Path))

		dto, err := k.uploadDto(r)
		if err != nil {
			k.log.Error("Failed to upload kpt file", sl.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.InternalServerError())
			return
		}

		err = k.UseCase.SaveKpt(ctx, &dto)
		if err != nil {
			if errs.Is(err, customErrors.ErrKptAlreadyExist) {
				k.log.Error("KPT already exists", sl.Err(err))
				w.WriteHeader(http.StatusConflict)
				render.JSON(w, r, response.Conflict("KPT already exists"))
				return
			}
			k.log.Error("Failed to save kpt", sl.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.InternalServerError())
			return
		}

		resp := Response{
			Response: response.OK(),
		}

		w.WriteHeader(http.StatusCreated)
		render.JSON(w, r, resp)
	}
}

func (k *SaveKptHandler) uploadDto(r *http.Request) (Dto.KptDto, error) {
	var err error
	var dto Dto.KptDto
	dto.Territory = &Dto.ExtractCadastralPlanTerritory{}

	kpt, kptHeaders, err := r.FormFile("kpt")
	if err != nil {
		k.log.Error("Failed to get kpt file", sl.Err(err))
		return dto, errors.Wrap(err, "Failed to get kpt file")
	}
	defer func(kpt multipart.File) {
		err := kpt.Close()
		if err != nil {
			k.log.Error("Failed to close save_kpt_handler file", sl.Err(err))
		}
	}(kpt)

	dto.KptHeaders = kptHeaders
	_, err = dto.Kpt.ReadFrom(kpt)
	if err != nil {
		k.log.Error("Failed to read kpt file", sl.Err(err))
		return dto, errors.Wrap(err, "Failed to read kpt file")
	}

	if err = xml.NewDecoder(bytes.NewReader(dto.Kpt.Bytes())).Decode(dto.Territory); err != nil {
		k.log.Error("Failed to parse kpt", sl.Err(err))
		return dto, errors.Wrap(err, "Failed to parse kpt")
	}

	return dto, nil
}
