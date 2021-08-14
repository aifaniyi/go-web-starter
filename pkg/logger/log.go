package logger

import (
	"log"
	"os"
	"strings"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

type Logger interface {
	Print(...interface{})
	Println(...interface{})
	Printf(string, ...interface{})
	Fatal(...interface{})
	Fatalln(...interface{})
}

var (
	Info  Logger
	Error Logger
	Warn  Logger
)

const (
	defaultLogFile = "/tmp/app.log"
)

func init() {
	switch {
	case strings.TrimSpace(os.Getenv("LOG_MODE")) == "STDOUT":
		Info = log.New(os.Stdout, "INFO: ", log.LstdFlags)
		Warn = log.New(os.Stdout, "WARN: ", log.LstdFlags)
		Error = log.New(os.Stdout, "ERROR: ", log.LstdFlags)

	default:
		logFile := strings.TrimSpace(os.Getenv("LOG_FILE"))
		if logFile == "" {
			logFile = defaultLogFile
		}

		rotate := &lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    100,
			MaxBackups: 30,
			MaxAge:     90,
			Compress:   true,
			LocalTime:  true,
		}

		Info = log.New(rotate, "INFO: ", log.LstdFlags)
		Error = log.New(rotate, "ERROR: ", log.LstdFlags)
		Warn = log.New(rotate, "WARN: ", log.LstdFlags)
	}
	Info.Println("logger started")
}
