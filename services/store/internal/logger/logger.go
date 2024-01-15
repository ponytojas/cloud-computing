package logger

import (
	"errors"
	"fmt"
	"os"
	"syscall"

	"go.uber.org/zap"
)

var sugar *zap.SugaredLogger

func init() {
	if os.Getenv("DEBUG") == "true" {
		logger, _ := zap.NewDevelopment()
		defer func(logger *zap.Logger) {
			err := logger.Sync()
			if err != nil && !errors.Is(err, syscall.ENOTTY) {
				fmt.Println(err)
			}
		}(logger)
		sugar = logger.Sugar()
		return
	} else {
		logger, _ := zap.NewProduction()
		defer func(logger *zap.Logger) {
			err := logger.Sync()
			if err != nil && !errors.Is(err, syscall.ENOTTY) {
				fmt.Println(err)
			}
		}(logger)
		sugar = logger.Sugar()
	}
}

func GetLogger() *zap.SugaredLogger {
	return sugar
}
