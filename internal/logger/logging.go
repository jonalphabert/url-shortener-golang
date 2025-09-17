package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

type LoggerType = logrus.Logger

func New() *LoggerType {
	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return l
}
