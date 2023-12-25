package logger

import (
	"github.com/google/logger"
	"os"
)

type Logger struct {
	*logger.Logger
}

func InitLogger(logPath string) (Logger, error) {
	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		return Logger{nil}, err
	}
	return Logger{logger.Init("LoggerExample", false, true, lf)}, nil
}
