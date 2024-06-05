package land_plots_handler

import (
	"context"
	"log/slog"
	"net/http"

	Dto "github.com/ShevelevEvgeniy/app/internal/dto"
	services "github.com/ShevelevEvgeniy/app/internal/service"
	"github.com/ShevelevEvgeniy/app/lib/api/response"
	"github.com/ShevelevEvgeniy/app/lib/logger/sl"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type LandPlotsHandler struct {
	service   services.LandPlotsService
	log       *slog.Logger
	validator *validator.Validate
}

func NewLandPlotsHandler(log *slog.Logger, lpService services.LandPlotsService, validator *validator.Validate) *LandPlotsHandler {
	lpHandler := &LandPlotsHandler{
		service:   lpService,
		log:       log,
		validator: validator,
	}

	return lpHandler
}

func (l *LandPlotsHandler) GetCoordinates(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.GetCoordinates.New"

		l.log = l.log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		l.log.Info("Received HTTP GET request", slog.String("path", r.URL.Path))

		var dto Dto.LandPlot

		queryParams := r.URL.Query()
		dto.CadNumber = queryParams.Get("cad_number")

		err := l.validator.Struct(dto)
		if err != nil {
			l.log.Error("Failed to validate request: ", sl.Err(err))
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, response.BadRequest("Data not found"))
			return
		}

		landPlotCoordinates, err := l.service.GetCoordinatesList(ctx, dto.CadNumber)
		if err != nil {
			l.log.Error("Failed to get coordinates", sl.Err(err))
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, response.InternalServerError())
			return
		}

		l.log.Info("request completed", slog.String("op", op))

		var res Response

		res.Coordinates = landPlotCoordinates.Coordinates
		res.CadNumber = landPlotCoordinates.CadNumber
		res.Status = response.OK()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, res)
	}
}
