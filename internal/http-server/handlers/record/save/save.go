package save

import (
    "net/http"
    "log/slog"
    "time"

    "github.com/go-playground/validator/v10"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/render"

    "github.com/solloball/aws_tg/internal/storage"
    "github.com/solloball/aws_tg/internal/logger/sl"
    "github.com/solloball/aws_tg/internal/lib/api/response"
    "github.com/solloball/aws_tg/internal/lib/random"
)

type Request struct {
    Record storage.Record `json:"record" validate:"required"`
    Alias string `json:"alias,omitempty"`
}

type Response struct {
    response.Response
    Alias string `json:"alias,omitempty"`
}

type RecordSaver interface {
    SaveRecord (rec storage.Record, alias string) (int64, error)
}

// TODO: move to config
const aliasLength = 10

func New(log *slog.Logger, recordSaver RecordSaver) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        const op = "handlers.url.save.New"

        log = log.With(
            slog.String("op", op),
            slog.String("request_id", middleware.GetReqID(r.Context())),
        )

        var req Request

        if err := render.DecodeJSON(r.Body, &req); err != nil {
            log.Error("failed to decode request body", sl.Err(err))

            render.JSON(w, r, response.Error("failed to decode request"))

            return 
        }
        
        log.Info("request body was decoded", slog.Any("request", req))

        if err := validator.New().Struct(req); err != nil {
            validatedErr := err.(validator.ValidationErrors)

            render.JSON(w, r, response.ValidationError(validatedErr))

            log.Error("invalid request", sl.Err(err))

            return
        }

        alias := req.Alias
        if alias == "" {
            // TODO: make checking for collisions
            alias = random.NewRandomString(aliasLength, time.Now().UnixNano())
        }

        id, err := recordSaver.SaveRecord(req.Record, alias)
        if err != nil {
            log.Info("failed to save url", sl.Err(err))
            
            render.JSON(w, r, response.Error("failed to save url"))

            return
        }

        log.Info("url added", slog.Int64("id", id))


        render.JSON(w, r, Response{
            Response: response.OK(),
            Alias: alias,
        })
    }
}
