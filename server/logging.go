package server

import (
	"context"
	"os"

	stdlog "log"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/term"
)

func InitializeLogging() {
	log.Logger = DefaultLogger()
	stdlog.SetFlags(0)
	stdlog.SetOutput(log.Logger)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = log.With().Caller().Logger()
	if term.IsTerminal(int(os.Stdout.Fd())) {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stderr,
				NoColor: true,
			},
		)
	}
}

func WithContext(ctx context.Context) zerolog.Context {
	ztx := log.With()
	if id, ok := ctx.Value("request_id").(string); ok {
		ztx = ztx.Str("request_id", id)
	}

	return ztx
}

func LoggerWithContext(ctx context.Context) zerolog.Logger {
	return WithContext(ctx).Logger()
}

func DefaultLogger() zerolog.Logger {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "host"
	}

	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	logger := log.With().
		Str("host", hostname).
		Str("version", "1.1").
		Logger()

	return logger
}
