package internal

import (
	"context"
	"log/slog"
	"os"
	"time"
)

// customHandler extends slog.JSONHandler to add extra fields
type customHandler struct {
	*slog.JSONHandler
}

func (h *customHandler) Handle(ctx context.Context, r slog.Record) error {
	// Add trace ID if available from Cloud Run request
	if traceHeader := ctx.Value("X-Cloud-Trace-Context"); traceHeader != nil {
		r.Add("logging.googleapis.com/trace", slog.StringValue(traceHeader.(string)))
	}

	return h.JSONHandler.Handle(ctx, r)
}

func SetupLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Convert time to RFC3339 format
			if a.Key == "time" {
				return slog.Attr{
					Key:   "timestamp",
					Value: slog.StringValue(a.Value.Time().Format(time.RFC3339)),
				}
			}
			return a
		},
	}

	handler := &customHandler{
		JSONHandler: slog.NewJSONHandler(os.Stdout, opts),
	}

	return slog.New(handler)
}
