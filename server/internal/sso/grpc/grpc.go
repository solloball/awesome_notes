package grpc

import (
    "log/slog"
    "context"
    "time"
    "fmt"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/credentials/insecure"
    grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
    grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"

    ssov1 "github.com/solloball/contract/gen/go/sso"
)

type Client struct {
    api ssov1.AuthClient
    log *slog.Logger
}

func New(
    ctx context.Context,
    log *slog.Logger,
    addr string,
    timeout time.Duration,
    retriesCount int,
) (*Client, error) {
    const op = "grpc.New"

    retrtyOptions := []grpcretry.CallOption {
        grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
        grpcretry.WithMax(uint(retriesCount)),
        grpcretry.WithPerRetryTimeout(timeout),
    }

    logOpt := []grpclog.Option {
        grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
    }

    // TODO: remake this
    cc, err := grpc.DialContext(ctx, addr,
        grpc.WithTransportCredentials(insecure.NewCredentials()),
        grpc.WithChainStreamInterceptor(
            grpclog.UnaryClientInterceptor(interceptorLogger(log), logOpt),
            grpcretry.UnaryClientInterceptor(retrtyOptions...),
        ),
    )
    if err != nil {
        return nil, fmt.Errorf("%s: %w", op, err)
    }

    return &Client {
        api: ssov1.NewAuthClient(cc),
        log: log,   
    }, nil
}

func (c *Client) IsAdmin(ctx context.Context, userID int64) (bool, error) {
    const op = "grpc.IsAdmin"

    resp, err := c.api.IsAdmin(ctx, &ssov1.IsAdminRequest{
        UserId: userID,
    })
    if err != nil {
        return false, fmt.Errorf("%s: %w", op, err)
    }

    return resp.IsAdmin, nil
}

func interceptorLogger(l *slog.Logger) grpclog.Logger {
    return grpclog.LoggerFunc(
        func(ctx context.Context, lvl grpclog.Level, msg string, fields ... any) {
            l.Log(ctx, slog.Level(lvl), msg, fields)
        },
    )
}

