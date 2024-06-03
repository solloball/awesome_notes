package main


import (
    "fmt"
    "log/slog"
    "log"
    "os"
    "errors"

    "github.com/solloball/aws_tg/internal/config"
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

    //TODO: init storage

    //TODO: init router

    //TODO: run server
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
