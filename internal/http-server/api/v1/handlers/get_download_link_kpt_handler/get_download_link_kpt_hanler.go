package get_download_link_kpt_handler

import (
	"context"
	Dto "github.com/ShevelevEvgeniy/app/internal/dto"
	"github.com/ShevelevEvgeniy/app/lib/api/response"
	"github.com/ShevelevEvgeniy/app/lib/logger/sl"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
)

type GetKptUseCaseInterface interface {
	GetLinkAndInfoKpt(ctx context.Context, log *slog.Logger, cadQuarter string) (Dto.KptInfo, error)
}

type GetDownloadLinkKptHandler struct {
	log       *slog.Logger
	useCase   GetKptUseCaseInterface
	validator *validator.Validate
}

func NewGetDownloadLinkKptHandler(log *slog.Logger, useCase GetKptUseCaseInterface, validator *validator.Validate) *GetDownloadLinkKptHandler {
	return &GetDownloadLinkKptHandler{
		log:       log,
		useCase:   useCase,
		validator: validator,
	}
}

func (h GetDownloadLinkKptHandler) GetDownloadLinkKpt(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.GetDownloadLinkKpt.New"

		h.log = h.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var dto Dto.KptDto

		queryParams := r.URL.Query()
		dto.CadQuarter = queryParams.Get("cad_quarter")

		if err := h.validator.Struct(dto); err != nil {
			h.log.Error("Failed to validate request: ", sl.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.BadRequest("Bad request"))
			return
		}

		kptInfo, err := h.useCase.GetLinkAndInfoKpt(ctx, h.log, dto.CadQuarter)
		if err != nil {
			h.log.Error("Failed to get kpt info", sl.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.InternalServerError())
			return
		}

		h.log.Info("request completed", slog.String("op", op))

		var res Response

		res.Response = response.OK()
		res.Result = &kptInfo

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, res)
	}
}
