package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

func Init() {
	// Set the output to stdout
	log.SetOutput(os.Stdout)

	// Set the log level
	log.SetLevel(logrus.DebugLevel)

	// Set the formatter with colored output
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}
