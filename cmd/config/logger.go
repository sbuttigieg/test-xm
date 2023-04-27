package config

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(f *os.File) (*logrus.Logger, error) {
	logLevel, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		return nil, err
	}

	return &logrus.Logger{
		Out:       io.MultiWriter(f, os.Stdout),
		Level:     logLevel,
		Formatter: &logrus.JSONFormatter{},
	}, nil
}
