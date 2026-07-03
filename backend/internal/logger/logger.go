package logger

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Logger struct {
	App   *log.Logger
	Error *log.Logger
}

func New(logDir string) (*Logger, func(), error) {
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, nil, err
	}

	appFile, err := NewDailyFileWriter(logDir, "app")
	if err != nil {
		return nil, nil, err
	}

	ginFile, err := NewDailyFileWriter(logDir, "gin")
	if err != nil {
		appFile.Close()
		return nil, nil, err
	}

	errorFile, err := NewDailyFileWriter(logDir, "error")
	if err != nil {
		appFile.Close()
		ginFile.Close()
		return nil, nil, err
	}

	appLogger := log.New(
		io.MultiWriter(os.Stdout, appFile),
		"[APP] ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	errorLogger := log.New(
		io.MultiWriter(os.Stderr, errorFile),
		"[ERROR] ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	gin.DefaultWriter = io.MultiWriter(os.Stdout, ginFile)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, errorFile)

	cleanup := func() {
		appFile.Close()
		ginFile.Close()
		errorFile.Close()
	}

	return &Logger{
		App:   appLogger,
		Error: errorLogger,
	}, cleanup, nil
}
