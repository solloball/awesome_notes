package save

import (
    "net/http"
    "log/slog"

    "github.com/go-chi/chi/v5/middleware"

    "github.com/solloball/aws_tg/internal/storage"
    "github.com/solloball/aws_tg/internal/lib/api/response"
)

type Request struct {
    Record storage.Record `json:"record" validate:"required,record"`
    Alias string `json:"alias,omitempty"`
}

type Response struct {
    response.Response
    ALias string `json:"alias,omitempty"`
}

type RecordSaver interface {
    SaveRecord (rec storage.Record, alias string) (int64, error)
}

func New(log *slog.Logger, recordSaver RecordSaver) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        const op = "handlers.url.save.New"

        log = log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )

        var req Request

        err := render.De
    }
}
