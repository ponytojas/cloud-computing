package logger

import (
	"os"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func init() {
	if os.Getenv("DEBUG") == "true" {
		logger, _ := zap.NewDevelopment()
		defer logger.Sync()
		sugar = logger.Sugar()
		return
	} else {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		sugar = logger.Sugar()
	}
}

func GetLogger() *zap.SugaredLogger {
	return sugar
}
