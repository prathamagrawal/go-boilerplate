package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	Level        logrus.Level
	Format       logrus.Formatter
	ReportCaller bool
	Color        bool
	Output       string
}

func NewLogger(config LoggerConfig) *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(config.Level)
	logger.SetFormatter(config.Format)
	logger.ReportCaller = config.ReportCaller

	if config.Output != "" {
		file, err := os.OpenFile(config.Output, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			logger.Fatal(err)
		}
		logger.SetOutput(file)
	}

	if config.Color {
		logger.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
			CallerPrettyfier: func(frame *runtime.Frame) (string, string) {
				filename := filepath.Base(frame.File)
				return filename, fmt.Sprintf("%s:%d", filename, frame.Line)
			},
		})
	}

	return logger
}

var DefaultLoggerConfig = LoggerConfig{
	Level:        logrus.InfoLevel,
	Format:       &logrus.TextFormatter{},
	ReportCaller: true,
	Color:        true,
	Output:       "",
}
