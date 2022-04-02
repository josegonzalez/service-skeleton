package server

import (
	"io"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/logger"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func getenv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}
	return value
}

func GetServer() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
		log.Debug().
			Str("http_method", httpMethod).
			Str("path", absolutePath).
			Str("handler_name", handlerName).
			Int("handler_count", nuHandlers).
			Msg("route")
	}

	middlewares := []gin.HandlerFunc{
		logger.SetLogger(
			logger.WithLogger(func(c *gin.Context, out io.Writer, latency time.Duration) zerolog.Logger {
				return LoggerWithContext(c)
			}),
		),
		requestid.New(),
		ContextRequestidMiddleware(),
	}
	r := gin.New()
	r.Use(middlewares...)
	return r
}

func StartServer(r *gin.Engine) {
	port, err := strconv.Atoi(getenv("PORT", "5000"))
	if err != nil {
		log.Error().Err(err).Msg("Overriding PORT with default port 5000")
		port = 5000
	}

	log.Debug().
		Int("port", port).
		Msg("listening")

	if err := r.Run(":" + strconv.Itoa(port)); err != nil {
		log.Fatal().Err(err).Send()
	}
}
