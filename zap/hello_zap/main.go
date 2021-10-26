package main

import (
	"time"

	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewDevelopment()
	logger.Info("Hello zap", zap.String("key", "value"), zap.Time("now", time.Now()))
}
