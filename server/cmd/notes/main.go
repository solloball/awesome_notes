package main


import (
    "fmt"
    "log/slog"
    "log"
    "os"
    "errors"
    "net/http"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/go-chi/cors"

    "github.com/solloball/aws_note/internal/config"
    "github.com/solloball/aws_note/internal/storage/sqlite"
    "github.com/solloball/aws_note/internal/http-server/handlers/record/save"
    "github.com/solloball/aws_note/internal/http-server/handlers/record/get"
    "github.com/solloball/aws_note/internal/logger/sl"
)

const (
    envLocal = "local"
    envDev = "dev"
    envProd = "prod"
)

func main() {
    conf := config.MustLoad()

    logger, err := setupLogger(conf.Env)
    if err != nil {
       log.Fatal(err) 
    }
    logger.Info("start notes", slog.String("env", conf.Env))
    logger.Debug("debug mode is enabled")

    storage, err := sqlite.New(conf.StoragePath)
    if err != nil {
        logger.Error("failed to init storage", sl.Err(err))
        os.Exit(1)
    }

    router := chi.NewRouter()
    router.Use(middleware.RequestID)
    router.Use(middleware.RealIP)
    // TODO:: refactor this
    router.Use(middleware.Logger)
    router.Use(middleware.Recoverer)
    router.Use(middleware.URLFormat)

    router.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"https://*", "http://*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: false,
        MaxAge:           300, // Maximum value not ignored by any of major browsers
    }))


    router.Post("/api/record", save.New(logger, storage))
    router.Get("/api/{alias}", get.New(logger, storage))

    // TODO: implement this
    //router.Delete("/record/{alias}, delete.New(logger, storage))

    logger.Info("starting server", slog.String("address", conf.Address))

    server := &http.Server{
        Addr: conf.Address,
        Handler: router,
        ReadTimeout: conf.Timeout,
        WriteTimeout: conf.Timeout,
        IdleTimeout: conf.IdleTimeout,
    }

    if err := server.ListenAndServe(); err != nil {
        logger.Error("failed to start server")
    }

    logger.Error("server stopped")
}

func setupLogger(env string) (*slog.Logger, error) {
    var log *slog.Logger
    
    switch  env {
    case envLocal:
        log = slog.New(
            slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}),
        )
    case envDev:
        log = slog.New(
            slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}),
        )
    case envProd:
        log = slog.New(
            slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelInfo}),
        )
    default:
        return nil, errors.New(fmt.Sprintf("Unknown type of env: %s", env))
    }

    return log, nil
}
