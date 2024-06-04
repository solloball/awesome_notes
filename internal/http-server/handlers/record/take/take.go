package take

import (
    "log/slog"
    "net/http"
    "errors"

    "github.com/solloball/aws_tg/internal/storage"
    "github.com/solloball/aws_tg/internal/lib/api/response"
    "github.com/solloball/aws_tg/internal/logger/sl"

    "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type RecordGetter interface {
    GetRecord(alias string) (storage.Record, error)
}

func New(log *slog.Logger, recordGetter RecordGetter) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        const op = "handlers.url.take.New"

        log = log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )

        alias := chi.URLParam(r, "alias")
		if alias == "" {
			log.Info("alias is empty")

			render.JSON(w, r, response.Error("invalid request"))

			return
		}

        resRecord, err := recordGetter.GetRecord(alias)
		if errors.Is(err, storage.ErrRecordNotFound) {
			log.Info("url not found", "alias", alias)

			render.JSON(w, r, response.Error("not found"))

			return
		}
        if err != nil {
			log.Error("failed to get url", sl.Err(err))

			render.JSON(w, r, response.Error("internal error"))

			return
		}


        log.Info(
            "got url",
            slog.String("author", resRecord.Author),
            slog.String("title", resRecord.Title),
            slog.String("note", resRecord.Note),
        )
    }
}
